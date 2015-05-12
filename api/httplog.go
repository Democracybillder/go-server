package hapi

import (
	"net/http"
	"time"

	"github.com/Democracybillder/go-server/lib/logger"
)

var log = logger.NewLog("BillHttp")

func (p *HttpApi) HTTPLogger(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		f(w, r)
		t2 := time.Now()
		log.Info.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
}
