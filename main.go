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

// getOrders gets and adds the top 20 (set by level=2)
// open orders for prodID and adds them to the files
// "./data/prodID/bids.csv" and "./data/prodID/asks.csv"
func getOrders(prodID string) {

	for {
		err := orderbook.AddToData(prodID, "2")
		if err != nil {
			log.Print(err)
		}
		time.Sleep(5 * time.Minute)	
	}
}

// getTrades gets and adds the trades for prodID
// to
// "./data/prodID/trades.csv".
// This happens every day. If there happen to be more than 1000
// trades in a day, this won't pick up all of them until the next day.
func getTrades(prodID string) {

	for {
		populateTrades(prodID)
		time.Sleep(24 * time.Hour)
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
