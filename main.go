package main

import (
	"time"
	
	"github.com/rogimus/coinjar/internal/orderbook"
)

func main() {

	for {
		go orderbook.AddToData("BTCAUD", "2")
		go orderbook.AddToData("ETHAUD", "2")
		time.Sleep(5 * time.Minute)
	}
}
