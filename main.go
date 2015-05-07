package main

import (
	"fmt"
	"net/http"

	"github.com/Democracybillder/go-server/api"
	"github.com/Democracybillder/go-server/billdb"
	"github.com/Democracybillder/go-server/lib/confer"
	"github.com/Democracybillder/go-server/lib/dbsql"
)

func main() {

	cn, err := confer.GetConf("config.json")
	if err != nil {
		fmt.Println("error configing:", err)
		return
	}

	db, err := dbsql.ConnectDB(&cn.Postgres)
	bdb := billdb.NewPostgres(db)
	httpApi := hapi.NewHttp(bdb)

	http.HandleFunc("/bills", httpApi.BillHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println(cn)
}
