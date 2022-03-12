## `trades.go` Documentation

`GetAllTrades` gets (at most) the last `limit` trades of 
the product `prodID` since time `after`.

`AddFromTime` adds (at most) the last `limit` trades of 
`prodID` since time `after` to `./data/prodID/trades.csv`.

`UpdateData` adds the first `limit` trades since n days to
`./data/prodID/trades.csv`. If `trades.csv` is not empty,
then it will add the first `limit` trades since
the date of the last entry,
or from `n` days ago (whichever is more recent).
The limit is limited to 1000 by the API.

`GetLastDate` returns the date of the last entry in `./data/prodID/trades.csv`
as a time.Time.
If the file is empty it returns Go's standard format time
(2006-01-02 15:04:05 -0700 MST).
If at any stage it exits from an error, it returns the time at exit.
