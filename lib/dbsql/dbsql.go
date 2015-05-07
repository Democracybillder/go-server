package dbsql

import (
	"database/sql"
	"fmt"

	"github.com/Democracybillder/go-server/lib/confer"
)

func ConnectDB(pc *confer.PostgresConf) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		pc.Host, pc.User, pc.Password, pc.DBname, pc.SSLmode)

	conn, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
		fmt.Println("db connection error")
	}

	return conn, nil
}
