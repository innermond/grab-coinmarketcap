package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var url_template string = "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing?start=%d&limit=%d&sortBy=market_cap&sortType=desc&convert=USD,BTC,ETH&cryptoType=all&tagType=all&audited=false&aux=ath,atl,high24h,low24h,num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,volume_7d,volume_30d"

func main() {
	// will collect json data
	f, err := os.Create(strconv.Itoa(int(time.Now().Unix())) + ".json")
	if err != nil {
		log.Fatal(err)
		remove(f)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	// paging parameters
	offset := 1
	limit := 100
	total_count := offset + 1
	loop_crt := 0
	// safe threshold to avoid infinite loop
	loop_max := 100

	for offset <= total_count {
		if loop_crt >= loop_max {
			log.Fatal("Aborted! Max loops overflowed")
			remove(f)
		}
		fmt.Printf("offset %d limit %d\n", offset, limit)
		datum, err := get_currencies(offset, limit)
		if err != nil {
			fmt.Println("could not get data")
			log.Fatal(err)
			remove(f)
		}

		//total_count = 102
		total_count, _ = strconv.Atoi(datum.TotalCount)
		offset += limit
		loop_crt++
		fmt.Println("loop ", loop_crt)

		for _, data := range datum.CryptoCurrencyList {
			fmt.Printf("%v\n", data)
			err = json.NewEncoder(w).Encode(data)
			if err != nil {
				log.Fatal(err)
				remove(f)
			}
		}
		err = w.Flush()
		if err != nil {
			log.Fatal(err)
			remove(f)
		}

		time.Sleep(3 * time.Second)
	}
}

func remove(f *os.File) error {
	abs, err := filepath.Abs("./" + f.Name())
	if err != nil {
		return err
	}
	return os.Remove(abs)
}

func get_currencies(offset int, limit int) (*Data, error) {
	url := fmt.Sprintf(url_template, offset, limit)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoded := CryptoResponse{}
	r := bufio.NewReader(res.Body)
	err = json.NewDecoder(r).Decode(&decoded)
	if err != nil {
		return nil, err
	}
	if decoded.Status.Error_code != "0" {
		return nil, errors.New(decoded.Status.Error_message)
	}

	return &decoded.Data, nil
}
