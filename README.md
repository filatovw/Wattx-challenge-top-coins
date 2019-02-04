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
                    request top_<limit> currencies  -> rank
                    response top_<limit> currencies <- rank
                    request prices for currencies   -> price
                    prices for given currencies     <- price
                aggregate responses
            <- return pricelist data
        <- return compiled response
        
What I'd add if I have time:
----------------------------

    - auto registrator service. It's already there, but I have some troubles with it.
    - small fast database with support of TTL tailored to `pricelist` service for cache
    - healthchecks across the infrastructure


How to start:
=============

Make sure you have docker, docker-compose and make installed.

Clone repository:

    git clone https://github.com/filatovw/Wattx-challenge-top-coins.git
    cd Wattx-challenge-top-coins 

Download images:

    make pull

Start infrastructure:

    make infra

Rollout configs:

    make configure

Start services:

    make restart

Watch logs:
    
    make logs


Get report in CSV:

    curl -X GET "http://0.0.0.0:8667/?limit=100" -H "content-type:text/csv"

Get report in JSON:

    curl -X GET "http://0.0.0.0:8667/?limit=100" -H "content-type:application/json"

or 

    curl -X GET "http://0.0.0.0:8667/?limit=100"