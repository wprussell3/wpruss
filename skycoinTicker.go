package main

import (
	"github.com/Kasmetski/cmcAPI"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"strconv"
	"time"
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
	getSky, err := cmcAPI.GetCoinInfo("skycoin")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(getSky.Symbol)
	fmt.Println(getSky.PriceUsd)
	fmt.Println(getSky.Two4HVolumeUsd)

	symbol := getSky.Symbol
	price := strconv.FormatFloat(getSky.PriceUsd, 'f', -1, 64)
	volume := strconv.FormatFloat(getSky.Two4HVolumeUsd, 'f', -1, 64)
	outputMessage := "Skycoin Info: " + ", Symbol: " + symbol + ", Price: " + price + ", 24 Hour USD Volume: " + volume

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

