package trades_test


import (
	"testing"
	"log"

	"github.com/rogimus/coinjar/internal/trades"
)

const timeLayout = "2006-01-02 15:04:05 -0700 MST"

func TestUpdateData (t *testing.T) {
	err := trades.UpdateData("BTCAUD", "2000")
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetAllOrders (t *testing.T) {
	
	//	prevTime := strconv.FormatInt(time.Now().Unix() - 30000, 10)
	_, err := trades.GetAllTrades("BTCAUD", "5", timeLayout)
	if err != nil {
		log.Fatal(err)
	}
	// for _, n := range trs {
	// 	fmt.Println(n)
	// }
}

func TestAddFromTime (t *testing.T) {

	err := trades.AddFromTime("BTCAUD", "10", timeLayout)
	if err != nil {
		log.Fatal(err)
	}
}

