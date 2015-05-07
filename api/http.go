package hapi

import (
	"encoding/json"
	"net/http"

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

func (p *HttpApi) BillHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.FormValue("state")
	term := r.FormValue("term")
	bill, _ := p.db.GetBills(state, term)

	js, err := json.Marshal(bill)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
