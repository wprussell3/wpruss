package main 

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"time"
	"strconv"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Coin struct {
ID               string  `json:"id"`
Name             string  `json:"name"`
Symbol           string  `json:"symbol"`
Rank             int     `json:"rank,string"`
PriceUsd         float64 `json:"price_usd,string"`
PriceBtc         float64 `json:"price_btc,string"`
Two4HVolumeUsd   float64 `json:"24h_volume_usd,string"`
MarketCapUsd     float64 `json:"market_cap_usd,string"`
AvailableSupply  float64 `json:"available_supply,string"`
TotalSupply      float64 `json:"total_supply,string"`
PercentChange1H  float64 `json:"percent_change_1h,string"`
PercentChange24H float64 `json:"percent_change_24h,string"`
PercentChange7D  float64 `json:"percent_change_7d,string"`
LastUpdated      string  `json:"last_updated"`
PriceCny		 string `json:"price_cny"`
DayVolumeCny     string `json:"24h_volume_cny"`
MarketCapCny	 string `json:"market_cap_cny"`
PriceEth		 string `json:"price_eth"`
DayVolumeEth     string `json:"24h_volume_eth"`
MarketCapEth	 string `json:"market_cap_eth"`
}

var (
	baseURL = "https://api.coinmarketcap.com/v1"
	url     string
)

func main() {
	ticker := time.NewTicker(time.Hour * 1)
	go func() {
	for t := range ticker.C {
	fmt.Println(t)

	twitterTicker()

}
}()
	time.Sleep(time.Hour * 24) //default stop set to 1 day from start time
    ticker.Stop()
    fmt.Println("Ticker stopped")
}

func twitterTicker(){
	getSkyCny, err := GetCoinInfo("skycoin/?convert=CNY")
	if err != nil {
		fmt.Println(err)
	}

	getSkyEth, err := GetCoinInfo("skycoin/?convert=ETH")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getSkyCny)
	fmt.Println(getSkyEth)

	priceUsd := strconv.FormatFloat(getSkyCny.PriceUsd, 'f', -1, 64)
	priceCny := getSkyCny.PriceCny
	priceBtc := strconv.FormatFloat(getSkyCny.PriceBtc, 'f', -1, 64)
	priceEth := getSkyEth.PriceEth
	dayVolume := strconv.FormatFloat(getSkyCny.Two4HVolumeUsd, 'f', -1, 64)

	outputMessage := "Skycoin -- Current Prices: " + priceUsd + "USD " + priceCny +"CNY " + priceBtc + "BTC " + priceEth + "ETH  - 24 Hour Volume: " + dayVolume + "USD"

	consumerKey := "KEY"
	consumerSecret := "SECRET"
	accessToken := "TOKEN"
	accessSecret := "SECRET"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}

	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's Name:%+v\n", user.ScreenName)

	tweet, resp, err := client.Statuses.Update(outputMessage, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tweet)
	fmt.Println(resp)

}

func GetCoinInfo(coin string) (Coin, error) {
	url = fmt.Sprintf("%s/ticker/%s", baseURL, coin)
	fmt.Println(url)
	resp, err := makeReq(url)
	if err != nil {
		log.Println(err)
		return Coin{}, err
	}
	var data []Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println(err)
		return Coin{}, err
	}

	return data[0], nil
}

func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func makeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := doReq(req)
	if err != nil {
		log.Println(err)
	}

	return resp, err
}
