package main

import (
	"time"
	"log"
	"fmt"
	
	"github.com/rogimus/coinjar/internal/orderbook"
	"github.com/rogimus/coinjar/internal/trades"
)

func main() {

	//go getOrders()
	//go getTrades()
	populateTrades("ETHAUD")

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
		err := trades.UpdateData("BTCAUD", "1000", 100)
		if err != nil {
			log.Print(err)
		}
		err = trades.UpdateData("ETHAUD", "1000", 100)
		time.Sleep(1 * time.Second)
	}
}

func populateTrades(prodID string) {

	for {
		prevDate,err := trades.GetLastDate(prodID)
		if err != nil {
			log.Print(err)
			return
		}
		if prevDate.Unix() < time.Now().Add(-24 * time.Hour).Unix() {
			trades.UpdateData(prodID, "1000", 100)
		} else {
			trades.UpdateData(prodID, "1000", time.Now().Add(-24 * time.Hour).Unix())
			return				
		}
	}
}
