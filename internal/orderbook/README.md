## Documentation

`GetAllOrders` gets a certain number of orders (determined by `level`)
for `prodID` and stores them in memory. See the 
[CoinJar API docs][https://docs.exchange.coinjar.com/data-api/#/introduction/security]
for more info on `level`. 

`AddToData` gets and stores a certain number of orders (determined by `level`)
for `prodID` and stores them in `./data/prodID/bids.csv` and
`./data/prodID/asks.csv`.
