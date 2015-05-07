package billdb

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Democracybillder/go-server/lib/confer"
	"github.com/Democracybillder/go-server/lib/dbsql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestDb struct {
	db *sql.DB
}

func (p *TestDb) TruncateBills() error {
	const query = "TRUNCATE TABLE bills CASCADE"
	_, err := p.db.Exec(query)
	return err
}

func (p *TestDb) InsertTestBillDesc(bill_id int, title string, state string, description string) error {
	const query = "INSERT INTO bills(bill_id, title, state, descr) VALUES ($1, $2, $3, $4)"
	_, err := p.db.Exec(query, bill_id, title, state, description)
	return err

}

func (p *TestDb) InsertTestBillLog(bill_id int, status_date time.Time, status int, last_action_date time.Time, last_action string) error {
	const query = "INSERT INTO bill_log(bill_id, status_date, status, last_action_date, last_action) VALUES ($1, $2, $3, $4, $5)"
	_, err := p.db.Exec(query, bill_id, status_date, status, last_action_date, last_action)
	return err

}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type BillTestSuite struct {
	suite.Suite
	testdb TestDb
	db     BillDb
}

// before each test
func (suite *BillTestSuite) SetupTest() {

	// clean DB
	err := suite.testdb.TruncateBills()
	assert.Empty(suite.T(), err)

	// Insert some test data
	err = suite.testdb.InsertTestBillDesc(testBill.Id, testBill.Title, testBill.State, testBill.Description)
	assert.Empty(suite.T(), err)

	err = suite.testdb.InsertTestBillDesc(testBill2.Id, testBill2.Title, testBill2.State, testBill2.Description)
	assert.Empty(suite.T(), err)

	err = suite.testdb.InsertTestBillLog(testBill.Id, testLog.StatusDate, testLog.Status, testLog.LastActionDate, testLog.LastAction)
	assert.Empty(suite.T(), err)

}

// test data
var t, _ = time.Parse("2006-Jan-02", "2006-Jan-02")

var testBill = BillDesc{
	Id:          987654321,
	Title:       "Supporting Kermit",
	State:       "CA",
	Description: "Really supporting Kermit",
}

var testBill2 = BillDesc{
	Id:          987654322,
	Title:       "Neglecting Bigbird",
	State:       "CA",
	Description: "Really Neglecting Bigbird",
}

var testLog = BillLog{
	StatusDate:     t,
	Status:         5,
	LastActionDate: t,
	LastAction:     "passed",
}

// Test bill fetching function with test data
func (suite *BillTestSuite) TestGetBills() {

	bl, err := suite.db.GetBills(testBill.State, "%supporting%")
	assert.Empty(suite.T(), err)

	assert.Equal(suite.T(), len(bl), 1, "database should return 1 item")
	assert.Equal(suite.T(), bl[0].Id, testBill.Id, "Id should be equal")
	assert.Equal(suite.T(), bl[0].Description.Title, testBill.Title, "Title should be equal")
	assert.Equal(suite.T(), bl[0].Description.Description, testBill.Description, "Description should be equal")
	assert.Equal(suite.T(), bl[0].Log[0].StatusDate, testLog.StatusDate, "Status date should be equal")
	assert.Equal(suite.T(), bl[0].Log[0].Status, testLog.Status, "Status should be equal")
	assert.Equal(suite.T(), bl[0].Log[0].LastActionDate, testLog.LastActionDate, "Last action date should be equal")
	assert.Equal(suite.T(), bl[0].Log[0].LastAction, testLog.LastAction, "Last action should be equal")

}

// Test bill fetching functionality when sending an empty term
func (suite *BillTestSuite) TestGetBillsEmptyTerm() {
	bl, err := suite.db.GetBills(testBill.State, "")
	assert.Empty(suite.T(), err)
	assert.Equal(suite.T(), len(bl), 2, "database should return a 2 item list")
}

// In order for 'go test' to run this suite, create
// a normal test function and pass our suite to suite.Run
func TestBillTestSuite(t *testing.T) {

	cn, err := confer.GetConf("config.json")
	if err != nil {
		fmt.Println("error configing:", err)
		return
	}

	db, err := dbsql.ConnectDB(&cn.Postgres)
	if err != nil {

		fmt.Printf("Error connecting to db: %v\n", err)
		return
	}

	billdb := NewPostgres(db)
	testdb := TestDb{db}
	bts := new(BillTestSuite)
	bts.db = billdb
	bts.testdb = testdb
	suite.Run(t, bts)
}
