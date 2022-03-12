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
	"net/url"
)


// Trades is a struct matching the json structure passed by the API.
type Trades []struct {
	Tid       int       `json:"tid"`
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Value     string    `json:"value"`
	TakerSide string    `json:"taker_side"`
	Timestamp time.Time `json:"timestamp"`
}

const timeLayout = "2006-01-02 15:04:05 -0700 MST"

// GetAllTrades gets (at most) the last /limit/ trades of /prodID/
// since time /after/.
func GetAllTrades (prodID, limit string, after time.Time) (Trades, error) {

	afterStr := strconv.FormatInt(after.Unix(), 10)

	URL, err := url.Parse("https://data.exchange.coinjar.com/products/id/trades?limit=limit&after=after")
	if err != nil {
		log.Fatal(err)
	}
	URL.Path =  fmt.Sprintf("/products/%s/trades", prodID)
	queries, err := url.ParseQuery(URL.RawQuery)
	if err != nil {
		log.Fatal(err)
	}
	queries.Set("limit", limit)
	queries.Set("after", afterStr)
	URL.RawQuery = queries.Encode()

	resp, err := http.Get(URL.String())
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

// AddFromTime adds (at most) the last /limit/ trades of /prodID/ since
// time /after/ to "./data/prodID/trades.csv".
func AddFromTime (prodID, limit string, after time.Time) error {
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
	currData, err := GetAllTrades(prodID, limit, after) 
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


// UpdateData adds the first /limit/ trades since n days to
// "./data/prodID/trades.csv". If trades.csv is not empty,
// then it will add the first /limit/ trades since
// the date of the last entry,
// or from n days ago (whichever is more recent).
// The limit is limited to 1000 by the API.
func UpdateData (prodID, limit string, n int64) error {
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

	prevTime, err := GetLastDate(prodID)
	if err != nil {
		log.Print(err)
		return err
	}
	prevTime = prevTime.Add(1 * time.Second)
	nDaysAgo := time.Now().Add(time.Duration(-n*24) * time.Hour)
	if prevTime.Unix() <= nDaysAgo.Unix() {
		prevTime = nDaysAgo
	}

	err = AddFromTime(prodID, limit, prevTime)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// GetLastDate returns the date of the last entry in ./data/prodID/trades.csv
// as a time.Time.
// If the file is empty it returns Go's standard format time
// (2006-01-02 15:04:05 -0700 MST).
// If at any stage it exits from an error, it returns the time at exit.
func GetLastDate(prodID string) (time.Time, error) {

 	baseDIR := fmt.Sprintf("/Users/roger/github.com/rogimus/coinjar/data/%s", prodID)
	for err := os.Mkdir(baseDIR, 0777); err != nil; {
 		if os.IsExist(err) {
 			break
 		} else {
 			return time.Now(),err
 		}
 	}
	
 	tradesDIR := fmt.Sprintf("%s/trades.csv", baseDIR)
 	tradesCSV, err := os.OpenFile(tradesDIR, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		tradesCSV.Close()
		return time.Now(),err
	}
	defer tradesCSV.Close()
	stat, err := os.Stat(tradesDIR)
	if err != nil {
		log.Print(err)
		return time.Now(),err
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
		return time.Now(),err
	}
	var PrevTime time.Time
	if len(lines) == 0 {
		PrevTime, err = time.Parse(timeLayout, timeLayout)
	} else {
		lastLine := strings.Split(lines[len(lines)-2], ",")
		PrevTime, err = time.Parse(timeLayout, lastLine[0])
		//  API requires seconds ^^^. Theoretically not a problem unless there
		// are more than 1000 trades in a single second.
		if err != nil {
			log.Fatal(err)
			return time.Now(),nil
		}
	}
	return PrevTime, nil
}
