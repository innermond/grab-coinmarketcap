# grab-coinmarketcap
## Experimental data extraction from api.coinmarketcap.com/data-api/v3/cryptocurrencies

It access https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing... and parse JSON into Golang structs, which one interested can *analyse*, transform, mangle, etc, and after it just saves all that data, unalterated, as a mere json text file.

Why this intermediary step before saving data? Because the main action should not be saving pristine data, but analysed/changed data. Which you are invited to do.

Please, while the code is functional, take it as BB (bare bone) and not as a BS(bulsh..) 
