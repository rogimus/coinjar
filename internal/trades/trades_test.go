package trades_test


import (
	"testing"
	"log"

	"github.com/rogimus/coinjar/internal/trades"
)

const timeLayout = "2006-01-02 15:04:05 -0700 MST"

func TestUpdateData (t *testing.T) {
	err := trades.UpdateData("BTCAUD", "1000")
	if err != nil {
		log.Fatal(err)
	}
	err = trades.UpdateData("ETHAUD", "1000") // the API is limited to 1000
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetAllOrders (t *testing.T) {
	
	_, err := trades.GetAllTrades("BTCAUD", "5", timeLayout)
	if err != nil {
		log.Fatal(err)
	}
}

func TestAddFromTime (t *testing.T) {

	err := trades.AddFromTime("BTCAUD", "10", timeLayout)
	if err != nil {
		log.Fatal(err)
	}
}

