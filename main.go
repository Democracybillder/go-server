package main

import (
	"net/http"

	"github.com/Democracybillder/go-server/api"
	"github.com/Democracybillder/go-server/billdb"
	"github.com/Democracybillder/go-server/lib/confer"
	"github.com/Democracybillder/go-server/lib/dbsql"
	"github.com/Democracybillder/go-server/lib/logger"
)

func main() {
	log := logger.NewLog("Bill-Server")

	cn, err := confer.GetConf("config.json")
	if err != nil {
		log.Error.Fatalf("Fetching configuration failed, please check config.json %v\n", err)
	}

	postgresConf := &cn.Postgres
	if postgresConf.DBname == "" {
		log.Error.Fatal("Postgres db name configuration not found")
	}

	if postgresConf.Host == "" {
		log.Error.Fatal("Postgres db host configuration not found")
	}

	if postgresConf.User == "" {
		log.Error.Fatal("Postgres db user configuration not found")
	}

	httpConf := &cn.BillServer
	if httpConf.Port == "" {
		log.Error.Fatal("BillServer port configuration not found")
	} else {
		log.Info.Printf("Listening on port %v\n", httpConf.Port)
	}

	log.Info.Println("Configuration successful")

	db, err := dbsql.ConnectDB(postgresConf)
	bdb := billdb.NewPostgres(db)
	httpApi := hapi.NewHttp(bdb)

	http.HandleFunc("/bills", httpApi.HTTPLogger(httpApi.BillHandler))

	err = http.ListenAndServe(":"+httpConf.Port, nil)
	if err != nil {
		log.Error.Fatalf("Error starting server %v\n", err)
	}
}
