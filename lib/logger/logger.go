package logger

import (
	"io/ioutil"
	"log"
	"os"
)

type Log struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func NewLog(pckg string) *Log {
	lg := &Log{}
	lg.Trace = log.New(ioutil.Discard,
		pckg+" TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	lg.Info = log.New(os.Stdout,
		pckg+" INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	lg.Warning = log.New(os.Stdout,
		pckg+" WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	lg.Error = log.New(os.Stderr,
		pckg+" ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	return lg
}
