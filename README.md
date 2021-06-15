# grab-coinmarketcap
## Experimental data extraction from api.coinmarketcap.com/data-api/v3/cryptocurrencies

It access https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing?start=%d&limit=%d&sortBy=market_cap&sortType=desc&convert=USD,BTC,ETH&cryptoType=all&tagType=all&audited=false&aux=ath,atl,high24h,low24    h,num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,volume_7d,volume_30d
and parse JSON into Golang structs, which one interested can analyse, transform, mangle, etc, and after it saves all that data, unalterated, as a mere json text file.

Please, while the code is functional, take it as BB (bare bone) and not as a BS(bulsh..) 
