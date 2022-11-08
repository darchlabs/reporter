package main

import (
	"fmt"
	"log"

	"github.com/darchlabs/reporter"
	"github.com/darchlabs/reporter/internal/commander"
	"github.com/darchlabs/reporter/internal/config"
	"github.com/darchlabs/reporter/internal/storage"
	reporterstorage "github.com/darchlabs/reporter/internal/storage/reporter"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	fmt.Println("Starting reporter")

	// read env values
	var conf config.Config

	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal("invalid env values")
	}

	// check serviceType is valid
	st := reporter.ServiceType(conf.ServiceType)

	// initialize storage
	s, err := storage.New(conf.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// initialize reporter storage
	reporterStorage := reporterstorage.New(s)

	// make process for getting service status (syncs, nodes or jobs) and save in database
	err = commander.Reporter(st, conf.ServiceURL, reporterStorage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done service_type=%s service_url=%s", conf.ServiceType, conf.ServiceURL)
}