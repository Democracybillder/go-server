// database layer package for Billder.
package billdb

import "time"

// properties of the BillDb object
type BillDb interface {
	GetBills(state string, term string) ([]*Bill, error)
}

type BillDesc struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	State       string `json:"state"`
	Description string `json:"description"`
}

type BillLog struct {
	StatusDate     time.Time `json:"status_date"`
	Status         int       `json:"status"`
	LastActionDate time.Time `json:"last_action_date"`
	LastAction     string    `json:"last_action"`
}

type Bill struct {
	Id          int        `json:"id"`
	Description BillDesc   `json:"description"`
	Log         []*BillLog `json:"log"`
}
