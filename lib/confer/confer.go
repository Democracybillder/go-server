package confer

import (
	"encoding/json"

	"github.com/Democracybillder/go-server/lib/logger"

	"os"
)

var log = logger.NewLog("Confer")

type Conf struct {
	Postgres   PostgresConf
	BillServer Server
}

type Server struct {
	Port string
}

type PostgresConf struct {
	Host     string
	User     string
	Password string
	DBname   string
	SSLmode  string
}

func GetConf(path string) (*Conf, error) {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	cnf := &Conf{}
	err := decoder.Decode(cnf)
	if err != nil {
		log.Error.Printf("Error decoding json config: %v\n", err)
		return nil, err
	}
	return cnf, nil
}
