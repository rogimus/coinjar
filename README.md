# CoinJar Exchange Scraper

An app developed in GO which scrapes CoinJar's Data RESTful API. 
Currently only gets the top 20 bids/asks every 5 minutes.

### TODO

- Track Trades
- Track Auctions
- Extend to their WebSocket-based API
- Containerise & move to cloud
- Add CLI 


## main.go Documentation

`getOrders` gets and adds the top 20 bids and asks every 5 minutes.

`getTrades` gets trades every day. Due to the way `populateTrades` 
handles trades within the last 24hrs, it is only reliably up to date >24s
in the past.

`populateTrades` gets all trades from 100 days ago until 24 hours ago
and also gets the first 1000 trades from the past 24 hours. Since trade history
persists, and since the purpose of this package (for now) is to get training
data for ML purposes (i.e. not for real-time trading) this is more than sufficient.
