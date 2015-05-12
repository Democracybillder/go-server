// database layer package for Billder.
package billdb

import (
	"database/sql"

	"github.com/Democracybillder/go-server/lib/logger"
	_ "github.com/bmizerany/pq"
)

var log = logger.NewLog("Bill-DB")

// Constructor method for BillDbPostgres, receives db connection and creates a db layer object
func NewPostgres(db *sql.DB) BillDb {
	return &BillDbPostgres{db}
}

type BillDbPostgres struct {
	db *sql.DB
}

// Returns a list of bill description dictionaries for a given term and US state.
func (p *BillDbPostgres) getBillDescByTerm(state string, term string) ([]*BillDesc, error) {
	const query = "SELECT bill_id, title, state, descr FROM bills WHERE lower(state) = $1 AND (lower(title) like $2 OR lower(descr) like $2 )"
	term = "%" + term + "%"
	rows, err := p.db.Query(query, state, term)
	if err != nil {
		log.Error.Printf("Error when retrieving bills by term: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	billdescriptions := []*BillDesc{}

	for rows.Next() {
		billdescription := &BillDesc{}
		if err = rows.Scan(&billdescription.Id, &billdescription.Title, &billdescription.State, &billdescription.Description); err != nil {

			log.Error.Printf("Error when scanning rows: %v\n", err)
			return nil, err
		}
		billdescriptions = append(billdescriptions, billdescription)

	}
	// Return an empty list if nothing returned by DB
	if billdescriptions == nil {
		billdescriptions = []*BillDesc{}
	}
	if err = rows.Err(); err != nil {
		log.Error.Printf("Error when closing rows: %v\n", err)
		return nil, err
	}
	return billdescriptions, nil
}

// Returns a list of all bill description dictionaries for given state.
func (p *BillDbPostgres) getBillDescByState(state string) ([]*BillDesc, error) {
	const query = "SELECT bill_id, title, state, descr FROM bills WHERE lower(state) = $1"

	rows, err := p.db.Query(query, state)
	if err != nil {
		log.Error.Printf("Error when retrieving bills by state: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	billdescriptions := []*BillDesc{}

	for rows.Next() {
		billdescription := &BillDesc{}
		if err = rows.Scan(&billdescription.Id, &billdescription.Title, &billdescription.State, &billdescription.Description); err != nil {

			log.Error.Printf("Error when scanning rows: %v\n", err)
			return nil, err
		}
		billdescriptions = append(billdescriptions, billdescription)

	}
	// Return an empty list if nothing returned by DB
	if billdescriptions == nil {
		billdescriptions = []*BillDesc{}
	}
	if err = rows.Err(); err != nil {
		log.Error.Printf("Error when closing rows: %v\n", err)
		return nil, err
	}
	return billdescriptions, nil
}

// Returns a list of all bill log dictionaries for given bill.
func (p *BillDbPostgres) getBillLog(billId int) ([]*BillLog, error) {
	const query = "SELECT status_date, status, last_action_date, last_action FROM bill_log WHERE bill_id = $1 ORDER BY last_action_date DESC"
	rows, err := p.db.Query(query, billId)
	if err != nil {
		log.Error.Printf("Error when retrieving bill logs: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	billLogs := []*BillLog{}

	for rows.Next() {
		billLog := &BillLog{}
		if err = rows.Scan(&billLog.StatusDate, &billLog.Status, &billLog.LastActionDate, &billLog.LastAction); err != nil {

			log.Error.Printf("Error when scanning rows: %v\n", err)
			return nil, err
		}
		billLogs = append(billLogs, billLog)

	}
	// Return an empty list if nothing returned by DB
	if billLogs == nil {
		billLogs = []*BillLog{}
	}

	if err = rows.Err(); err != nil {
		log.Error.Printf("Error when closing rows: %v\n", err)
		return nil, err
	}

	return billLogs, nil
}

// Returns descriptions and logs for each bill
// If term provided, calls getBillLog & getBillDescByTerm,
// otherwise, calls getBillLog & getBillDescByState.
func (p *BillDbPostgres) GetBills(state string, term string) ([]*Bill, error) {
	bills := []*Bill{}
	billDescs := []*BillDesc{}

	if term == "" {
		billDescs, _ = p.getBillDescByState(state)
	} else {
		billDescs, _ = p.getBillDescByTerm(state, term)
	}

	for _, bl := range billDescs {
		bill := &Bill{}

		bill.Id = bl.Id
		bill.Description.Title = bl.Title
		bill.Description.State = bl.State
		bill.Description.Description = bl.Description
		bill.Log, _ = p.getBillLog(bill.Id)

		bills = append(bills, bill)
	}
	return bills, nil
}
