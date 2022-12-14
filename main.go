package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/andreaciri/smart-grid-pi/led"
	"github.com/andreaciri/smart-grid-pi/relay"
	"github.com/andreaciri/smart-grid-pi/solaredge"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("starting smart-grid-pi...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error reading .env")
	}

	refreshTimeSeconds, err := strconv.Atoi(os.Getenv("REFRESH_TIME_SECONDS"))
	if err != nil {
		log.Fatal("error invalid REFRESH_TIME_SECONDS")
	}

	minimumPowerSurplusWatt, err := strconv.Atoi(os.Getenv("MIN_POWER_SURPLUS_WATT"))
	if err != nil {
		log.Fatal("error invalid MIN_POWER_SURPLUS_WATT")
	}

	solarEdgeClient := solaredge.NewClient(
		os.Getenv("SOLAREDGE_API_BASE_URL"),
		os.Getenv("SOLAREDGE_SITE_ID"),
		os.Getenv("SOLAREDGE_API_KEY"),
	)

	relay, err := relay.NewRelay()
	if err != nil {
		log.Fatal("relay error: ", err.Error())
	}

	led, err := led.NewLed()
	if err != nil {
		log.Fatal("led error: ", err.Error())
	}

	service := NewService(
		solarEdgeClient,
		relay,
		led,
		time.Duration(refreshTimeSeconds)*time.Second,
		minimumPowerSurplusWatt,
	)

	err = service.Run()

	log.Fatal("service error: ", err.Error())
}
