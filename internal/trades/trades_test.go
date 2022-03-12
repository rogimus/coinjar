package trades_test


import (
	"testing"
	"log"
	"time"

	"github.com/rogimus/coinjar/internal/trades"
)

const timeLayout = "2006-01-02 15:04:05 -0700 MST"

func TestUpdateData (t *testing.T) {
	// err := trades.UpdateData("BTCAUD", "1000")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err := trades.UpdateData("ETHAUD", "1000", 100) // the API is limited to 1000
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetAllOrders (t *testing.T) {

	after,_ := time.Parse(timeLayout,timeLayout)
	_, err := trades.GetAllTrades("BTCAUD", "5", after)
	if err != nil {
		log.Fatal(err)
	}
}

func TestAddFromTime (t *testing.T) {

	after,_ := time.Parse(timeLayout,timeLayout)
	err := trades.AddFromTime("BTCAUD", "10", after)
	if err != nil {
		log.Fatal(err)
	}
}

