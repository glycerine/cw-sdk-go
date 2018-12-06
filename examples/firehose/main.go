package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"code.cryptowat.ch/ws-client-go"
	"code.cryptowat.ch/ws-client-go/examples/kraken-trades/cwrest"
)

const (
	exchangeSymbol = "kraken"
)

var (
	apiKey    = flag.String("apikey", "", "API key to use.")
	secretKey = flag.String("secretkey", "", "Secret key to use.")
)

func init() {
	flag.Parse()
}

func main() {
	rest := cwrest.NewCWRESTClient("https://api.cryptowat.ch")

	// Get exchange description, in particular we'll need the ID to use it
	// in stream subscriptions

	// Get market descriptions, to know symbols (like "btcusd") by integer ID
	marketsSlice, err := rest.GetMarketsIndex()
	if err != nil {
		log.Fatalf("failed to get markets of %s: %s", exchangeSymbol, err)
	}

	markets := map[int]cwrest.MarketDescr{}
	for _, market := range marketsSlice {
		markets[market.ID] = market
	}

	// Create a new stream connection instance. Note that the actual connection
	// will happen later.
	c, err := wsclient.NewStreamClient(&wsclient.WSParams{
		URL: "wss://stream.cryptowat.ch",

		Subscriptions: []string{
			"markets:*:trades",
		},

		APIKey:    *apiKey,
		SecretKey: *secretKey,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	var lastError error

	c.OnError(func(err error, disconnecting bool) {
		if disconnecting {
			lastError = err
			return
		}

		log.Printf("Error: %s", err.Error())
	})

	// Ask for the state transition updates, and present them to the user somehow
	c.OnStateChange(
		wsclient.ConnStateAny,
		func(oldState, state wsclient.ConnState) {
			causeStr := ""
			if lastError != nil {
				causeStr = fmt.Sprintf(" (%s)", lastError)
				lastError = nil
			}
			log.Printf(
				"State updated: %s -> %s%s",
				wsclient.ConnStateNames[oldState],
				wsclient.ConnStateNames[state],
				causeStr,
			)
		},
	)

	// Listen for received market messages and print them
	c.OnMarketData(func(market wsclient.Market, md wsclient.MarketData) {
		if md.TradesUpdate != nil {
			tradesUpdate := md.TradesUpdate
			for _, trade := range tradesUpdate.Trades {
				fmt.Printf(
					"%v %34s %16s %16s\n",
					trade.Timestamp, market.ExchangeID+" "+strings.ToUpper(market.CurrencyPairID), trade.Price, trade.Amount,
				)
			}
		}
	})

	// Finally, connect.
	c.Connect()

	// Setup OS signal handler
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Wait until the OS signal is received, at which point we'll close the
	// connection and quit
	<-interrupt
	log.Printf("Closing connection...\n")

	if err := c.Close(); err != nil {
		log.Fatalf("Failed to close connection: %s", err)
	}
}
