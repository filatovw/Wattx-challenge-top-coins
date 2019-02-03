# Wattx-challenge-top-coins

Full description of a challenge is available [here](https://github.com/WATTx/code-challenges/blob/master/software-engineer-challenge-top-coins.md)


Architecture
============

Full solution consists of several services:

- API
- pricelist
- rank
- price

`API` service accept `HTTP` requests with `json` and `csv` headers and provides corresponding responces.
`pricelist` serves `gRPC` calls from `API`, this also rules cached results. It has its own database with prepared results.
`rank` service drains data from `https://min-api.cryptocompare.com/`. This service contains all the data about ranks.
`price` service drains data from `https://coinmarketcap.com/api/`. This service contains all the data about prices.

    Client request -> 
        API -> 
            pricelist ->
                if data still fresh -> return response
                else:
                    request top_<limit> currencies  -> rank
                    response top_<limit>            <- rank
                    request prices for currencies   -> price
                    prices for given currencies     <- price
                aggregate responses
                store into local database
            <- return pricelist data
        <- return compiled response
        
Both `rank` and `price` services drain actual info from the external APIs regularely (similar to cron task).

TODO:

- [x] `rank` service functionality
- [ ] `price` service functionality
- [x] skeleton for the API HTTP server
- [x] support CSV/JSON formats
- [x] protofiles
- [ ] protofiles compilator
- [ ] gRPC servers & clients
- [ ] `pricelist` aggregation
- [ ] `pricelist` storage
- [x] API service `/healthcheck` 
- [ ] API service `/` with support of `limit` parameter
- [ ] `price` storage
- [ ] `rank` storage
- [ ] `pricelist` storage
- [ ] explanation for development process
- [x] introduce service discovery
- [x] introduce config management

What I'd add if I have time:
----------------------------

    - distributed cron service that should set tasks in a queue
    - `rank` and `price` services refresh its data from APIs by message from the corresponding queue
    - there should be admin panel for a safe operational control

