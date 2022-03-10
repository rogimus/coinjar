package main

import (
	"time"
	
	"github.com/rogimus/coinjar/orderbook"
)

func main() {

	for {
		orderbook.AddToData("BTCAUD", "2")
		orderbook.AddToData("ETHAUD", "2")
		time.Sleep(5 * time.Minute)
	}
}
