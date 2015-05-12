package hapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Democracybillder/go-server/billdb"
)

func NewHttp(db billdb.BillDb) *HttpApi {

	return &HttpApi{
		db: db,
	}
}

type HttpApi struct {
	db billdb.BillDb
}

func (p *HttpApi) parseUrlParams(r *http.Request, params []string) map[string]string {
	parsedParams := map[string]string{}
	r.ParseForm()
	for _, param := range params {
		parsedParam := r.FormValue(param)
		parsedParam = strings.ToLower(strings.TrimSpace(parsedParam))
		parsedParams[param] = parsedParam

	}
	return parsedParams
}

func (p *HttpApi) BillHandler(w http.ResponseWriter, r *http.Request) {
	parsed := p.parseUrlParams(r, []string{"state", "term"})

	if parsed["state"] == "" {
		http.Error(w, "Error: Missing param: state ", http.StatusBadRequest)
		log.Error.Printf("Error: Missing param: state [%s] %q ", r.Method, r.URL.String())
	} else {

		bill, _ := p.db.GetBills(parsed["state"], parsed["term"])

		js, err := json.Marshal(bill)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Error.Printf("error with bill response [%s] %q ", r.Method, r.URL.String())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}
