package main

type CryptoResponse struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	CryptoCurrencyList []CryptoCurrency `json:"cryptoCurrencyList"`
	TotalCount         string           `json:"totalCount"`
}
type Status struct {
	Timestamp     string  `json:"timestamp"`
	Error_code    string  `json:"error_code"`
	Error_message string  `json:"error_message"`
	Elapsed       string  `json:"elapsed"`
	Credit_count  float64 `json:"credit_count"`
}

type CryptoCurrency struct {
	ID                float64  `json:"id"`
	Name              string   `json:"name"`
	Symbol            string   `json:"symbol"`
	Slug              string   `json:"slug"`
	Tags              []string `json:"tags"`
	CmcRank           float64  `json:"cmcRank"`
	MarketPairCount   float64  `json:"marketPairCount"`
	CirculatingSupply float64  `json:"circulatingSupply"`
	TotalSupply       float64  `json:"totalSupply"`
	MaxSupply         float64  `json:"maxSupply"`
	Ath               float64  `json:"ath"`
	Atl               float64  `json:"atl"`
	High24h           float64  `json:"high24h"`
	Low24h            float64  `json:"low24h"`
	IsActive          float64  `json:"isActive"`
	LastUpdated       string   `json:"lastUpdated"`
	DateAdded         string   `json:"dateAdded"`
	Quotes            []Quote  `json:"quotes"`
}

type Quote struct {
	Name                     string   `json:"name"`
	Price                    float64  `json:"price"`
	Volume24h                float64  `json:"volume24h"`
	Volume7d                 float64  `json:"volume7d"`
	Volume30d                float64  `json:"volume30d"`
	MarketCap                float64  `json:"marketCap"`
	PercentChange1h          float64  `json:"percentChange1h"`
	PercentChange24h         float64  `json:"percentChange24h"`
	PercentChange7d          float64  `json:"percentChange7d"`
	LastUpdated              string   `json:"lastUpdated"`
	PercentChange30d         float64  `json:"percentChange30d"`
	PercentCchange60d        float64  `json:"percentChange60d"`
	PercentCchange90d        float64  `json:"percentChange90d"`
	FullyDilluttedMarketCap  float64  `json:"fullyDilluttedMarketCap"`
	Dominance                float64  `json:"dominance"`
	Tvl                      float64  `json:"tvl"`
	Turnover                 float64  `json:"turnover"`
	YtdPriceChangePercentage float64  `json:"ytdPriceChangePercentage"`
	Platform                 Platform `json:"platform"`
	IsAudited                bool     `json:"isAudited"`
}

type Platform struct {
	ID           float64 `json:"id"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	TokenAddress string  `json:"token_address"`
	Slug         string  `json:"slug"`
}
