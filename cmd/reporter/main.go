package main

import (
	"fmt"
	"log"
	"os"

	"github.com/darchlabs/reporter"
	"github.com/darchlabs/reporter/internal/commander"
	"github.com/darchlabs/reporter/internal/storage"
	reporterstorage "github.com/darchlabs/reporter/internal/storage/reporter"
)

func main() {
	// get DATABASE_URL
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("invalid DATABASE_URL environment value")
	}

	// get SERVICE_URL
	serviceURL := os.Getenv("SERVICE_URL")
	if serviceURL == "" {
		log.Fatal("invalid SERVICE_URL environment value")
	}

	// get SERVICE_TYPE
	serviceType := os.Getenv("SERVICE_TYPE")
	if serviceType == "" {
		log.Fatal("invalid SERVICE_TYPE environment value")
	}

	// check serviceType is valid
	st := reporter.ServiceType(serviceType)

	// initialize storage
	s, err := storage.New(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// initialize reporter storage
	reporterStorage := reporterstorage.New(s)

	// make process for getting service status (syncs, nodes or jobs) and save in database
	err = commander.Reporter(st, serviceURL, reporterStorage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done service_type=%s service_url=%s", serviceType, serviceURL)
	

	// fmt.Println("BEFORE GET")
	// gr, err := reporterStorage.GetGroupReport(groupReport.CreatedAt.Unix(), reporter.ServiceTypeSynchronizer)
	// if err != nil {
	// 	fmt.Println(22222)
	// 	log.Fatal(err)
	// }

	// fmt.Println("AFTER")
	// fmt.Println("AFTER")
	// fmt.Println("AFTER", gr)

	// list, err := reporterStorage.ListGroupReports()
	// 	if err != nil {
	// 	fmt.Println(22222)
	// 	log.Fatal(err)
	// }

	// fmt.Println("ALLAALALALA")
	// fmt.Println("ALLAALALALA")
	// fmt.Println("ALLAALALALA")
	// fmt.Println(list)

	// for _, v := range list {
	// 	fmt.Println(5555555)
	// 	fmt.Printf("%+v", v)
	// }

	// fetch to jobs

	// fetch to nodes

	// initilize struct
	// - add array with len of events
	// for i, e := 
	// groupReport := &reporter.GroupReport{
	// 	Synchronizers: &reporter.Report{

	// 	},
	// }
	
	// insert in dabase
	// err = reporterStorage.InsertGroupReport(groupReport)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("inserted group report timestamp=%v", groupReport.CreatedAt.Unix())
}