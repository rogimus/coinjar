package orderbook

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"log"
)

const baseURL = "https://data.exchange.coinjar.com/products/"

type Orders struct {
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

func GetAllOrders (prodID string, level string) (Orders, error) {

	URL := baseURL + prodID + "/book?level=" + level
	resp, err := http.Get(URL)
	if err != nil {
		log.Print(err)
		resp.Body.Close()
		return Orders{}, err
	} 
	defer resp.Body.Close()

	var response Orders
	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err
	
}

func AddToData (prodID string, level string) error {

	baseDIR := fmt.Sprintf("./data/%s", prodID)
	for err := os.Mkdir(baseDIR, 0777); err != nil; {
		if os.IsExist(err) {
			break
		} else {
			return err
		}
	}
	
	currTime := time.Now().String()
	currData, err := GetAllOrders(prodID, level)
	if err != nil {
		return err
	}

	// add the bids data

	bidsDIR := fmt.Sprintf("%s/bids.csv", baseDIR)
	bidsCSV, err := os.OpenFile(bidsDIR, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		bidsCSV.Close()
		return err
	}

	for _, bid := range currData.Bids {
		t := fmt.Sprintf("%s,%s,%s\n", currTime, bid[0], bid[1])
		if _, err := bidsCSV.WriteString(t); err != nil {
			log.Println(err)
			return err
		}
	}
	bidsCSV.Close()

	
	// add the asks data
	asksDIR := fmt.Sprintf("%s/asks.csv", baseDIR)
	asksCSV, err := os.OpenFile(asksDIR, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		asksCSV.Close()
		return err
	}

	for _, ask := range currData.Asks {
		t := fmt.Sprintf("%s,%s,%s\n", currTime, ask[0], ask[1])
		if _, err := asksCSV.WriteString(t); err != nil {
			log.Println(err)
			return err
		}
	}
	asksCSV.Close()

	return nil
}
