package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/andreaciri/smart-grid-pi/relay"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("starting smart-grid-pi...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error reading .env")
	}

	seconds, err := strconv.Atoi(os.Getenv("TICKER_PERIOD_SECONDS"))
	if err != nil {
		log.Fatal("error invalid TICKER_PERIOD_SECONDS")
	}
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	errChan := make(chan error)

	// solarEdgeClient := solaredge.NewClient(
	// 	os.Getenv("SOLAREDGE_API_BASE_URL"),
	// 	os.Getenv("SOLAREDGE_SITE_ID"),
	// 	os.Getenv("SOLAREDGE_API_KEY"),
	// )

	relay, err := relay.NewRelay()
	if err != nil {
		log.Fatal("relay error: ", err.Error())
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				// power, err := solarEdgeClient.GetCurrentPower()
				// if err != nil {
				// 	errChan <- err
				// }
				// log.Println("Power! ", power)

				relay.Toggle()
			}
		}
	}()

	err = <-errChan
	ticker.Stop()
	log.Fatal("error: ", err.Error())
}
