package confer

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conf struct {
	Postgres PostgresConf
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
		fmt.Println("error:", err)
		return nil, err
	}
	return cnf, nil
}
