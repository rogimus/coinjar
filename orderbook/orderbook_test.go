package orderbook_test

import (
	"testing"

	"./orderbook"
)

func TestGetAllOrders (t *testing.T) {
	orderbook.GetAllOrders("BTCAUD", "1")
}

func TestAddToData (t *testing.T) {
	orderbook.AddToData("BTCAUD", "1")
}
