// Copyright 2022 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"

	"github.com/veraison/services/config"
	"github.com/veraison/services/kvstore"
	"github.com/veraison/services/policy"
	"github.com/veraison/services/vts/pluginmanager"
	"github.com/veraison/services/vts/policymanager"
	"github.com/veraison/services/vts/trustedservices"
)

func main() {
	v, err := config.ReadRawConfig("", false)
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	subs, err := config.GetSubs(v, "ta-store", "en-store", "po-store",
		"*po-agent", "plugin", "*vts")
	if err != nil {
		log.Fatal(err)
	}

	taStore, err := kvstore.New(subs["ta-store"])
	if err != nil {
		log.Fatalf("trust anchor store initialisation failed: %v", err)
	}

	enStore, err := kvstore.New(subs["en-store"])
	if err != nil {
		log.Fatalf("endorsement store initialization failed: %v", err)
	}

	poStore, err := policy.NewStore(subs["po-store"])
	if err != nil {
		log.Fatalf("policy store initialization failed: %v", err)
	}

	policyManager, err := policymanager.New(subs["po-agent"], poStore)
	if err != nil {
		log.Fatalf("policy manager initialization failed: %v", err)
	}

	pluginManager := pluginmanager.New()
	if err := pluginManager.Init(subs["plugin"]); err != nil {
		log.Fatalf("plugin manager initialization failed: %v", err)
	}

	// from this point onwards taStore, enStore and pluginManager are owned by vts
	vts := trustedservices.NewGRPC(taStore, enStore, pluginManager, policyManager)

	if err = vts.Init(subs["vts"]); err != nil {
		log.Fatalf("VTS initialisation failed: %v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go vtsRun(vts, done)
	go sigWaiter(sigs, done)

	<-done

	log.Println("stopping VTS")
	if err := vts.Close(); err != nil {
		log.Println("VTS termination failed:", err)
	}
	log.Println("bye!")
}

func vtsRun(vts trustedservices.ITrustedServices, done chan bool) {
	if err := vts.Run(); err != nil {
		log.Println("VTS failed:", err)
	}

	done <- true
}

func sigWaiter(sigs chan os.Signal, done chan bool) {
	sig := <-sigs

	log.Println(sig, "received, exiting")

	done <- true
}
