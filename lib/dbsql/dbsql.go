package dbsql

import (
	"database/sql"
	"fmt"

	"github.com/Democracybillder/go-server/lib/confer"
	"github.com/Democracybillder/go-server/lib/logger"
)

func ConnectDB(pc *confer.PostgresConf) (*sql.DB, error) {

	log := logger.NewLog("SQL")

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		pc.Host, pc.User, pc.Password, pc.DBname, pc.SSLmode)

	conn, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Error.Fatalf("Unable to connect to database %v\n", err)
	} else {
		log.Info.Printf("Connected to database %v\n", pc.DBname)
	}
	return conn, nil
}
