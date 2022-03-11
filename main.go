package main

import (
	"time"
	"log"
	"fmt"
	
	"github.com/rogimus/coinjar/internal/orderbook"
	"github.com/rogimus/coinjar/internal/trades"
)

func main() {

	go getOrders()
	go getTrades()

	fmt.Scanln()
	
}

func getOrders() {

	for {
		err := orderbook.AddToData("BTCAUD", "2")
		if err != nil {
			log.Print(err)
		}
		err = orderbook.AddToData("ETHAUD", "2")
		if err != nil {
			log.Print(err)
		}
		time.Sleep(5 * time.Minute)	
	}
}

func getTrades() {

	for {
		err := trades.UpdateData("BTCAUD", "1000")
		if err != nil {
			log.Print(err)
		}
		err = trades.UpdateData("ETHAUD", "1000")
		time.Sleep(1 * time.Second)
	}
}
