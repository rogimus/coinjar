package trades 

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"io"
	"time"
	"strings"
	"strconv"
)

const baseURL = "https://data.exchange.coinjar.com/products/"

type Trades []struct {
	Tid       int       `json:"tid"`
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Value     string    `json:"value"`
	TakerSide string    `json:"taker_side"`
	Timestamp time.Time `json:"timestamp"`
}

const timeLayout = "2006-01-02 15:04:05 -0700 MST"


func GetAllTrades (prodID, limit, after string) (Trades, error) {

	prevTime, err := time.Parse(timeLayout, after)
	if err != nil {
		log.Print(err)
		return Trades{}, err
	}
	prevTimeStr := strconv.FormatInt(prevTime.Unix(), 10)

	URL := fmt.Sprintf("%s%s/trades?limit=%s&after=%s", baseURL, prodID, limit, prevTimeStr)
	resp, err := http.Get(URL)
	if err != nil {
		log.Print(err)
		resp.Body.Close()
		return Trades{}, err
	} 
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return Trades{}, err
	}

	var response Trades
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Print(err)
		return Trades{}, err
	}
	return response, err
}


func AddFromTime (prodID, limit, after string) error {
	baseDIR := fmt.Sprintf("/Users/roger/github.com/rogimus/coinjar/data/%s", prodID)
	for err := os.Mkdir(baseDIR, 0777); err != nil; {
 		if os.IsExist(err) {
 			break
 		} else {
 			return err
 		}
 	}
	
 	tradesDIR := fmt.Sprintf("%s/trades.csv", baseDIR)
 	tradesCSV, err := os.OpenFile(tradesDIR, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		tradesCSV.Close()
		return err
	}
	defer tradesCSV.Close()
	currData, err := GetAllTrades(prodID, limit,  after) 
 	if err != nil {
		log.Print(err)
 		return err
 	}

	// CSV Format:
	//
	// TIME, TID, TAKER_SIDE, SIZE, PRICE, VALUE
	//
	// Struct format:
	//
	// type Trades []struct {
	// 	Tid       int       `json:"tid"`
	// 	Price     string    `json:"price"`
	// 	Size      string    `json:"size"`
	// 	Value     string    `json:"value"`
	// 	TakerSide string    `json:"taker_side"`
	// 	Timestamp time.Time `json:"timestamp"` }
	//
	for _, trade := range currData {
		t := fmt.Sprintf("%s,%d,%s,%s,%s,%s\n", trade.Timestamp, trade.Tid, trade.TakerSide, trade.Size, trade.Price, trade.Value)
		if _, err := tradesCSV.WriteString(t); err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}


func UpdateData (prodID, limit string) error {
 	baseDIR := fmt.Sprintf("/Users/roger/github.com/rogimus/coinjar/data/%s", prodID)
	for err := os.Mkdir(baseDIR, 0777); err != nil; {
 		if os.IsExist(err) {
 			break
 		} else {
 			return err
 		}
 	}
	
 	tradesDIR := fmt.Sprintf("%s/trades.csv", baseDIR)
 	tradesCSV, err := os.OpenFile(tradesDIR, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		tradesCSV.Close()
		return err
	}
	defer tradesCSV.Close()
	stat, err := os.Stat(tradesDIR)
	if err != nil {
		log.Print(err)
		return err
	}
	length := stat.Size()
	var res []byte
	var start int64
	var lines []string
	if length == 0 {
		lines = make([]string, 0)
		err = nil
	} else if 0 > length-400 {
		start = 0
		res = make([]byte, length)
		_, err = tradesCSV.Read(res)
		lines = strings.Split(string(res), "\n")
	} else {
		start = length-400
		res = make([]byte, 400)
		_, err = tradesCSV.ReadAt(res, start)
		lines = strings.Split(string(res), "\n")
	}
	if err != nil {
		log.Print(err)
		return err
	}
	var prevTime string
		
	if len(lines) == 0 {
		unformatedPrevTime := time.Now().Add(-144000 * time.Minute)
		prevTime = unformatedPrevTime.Format(timeLayout)
	} else {
		lastLine := strings.Split(lines[len(lines)-2], ",")
		temp, err := time.Parse(timeLayout, lastLine[0])
		prevTime = temp.Add(1 * time.Second).String()
		//  API requires seconds ^^^. Theoretically not a problem unless there
		// are more than 1000 trades in a single second.
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}

	return AddFromTime(prodID, limit, prevTime)

}
