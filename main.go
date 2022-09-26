package main

import (
	"log"
	"os"
	"strconv"
	"time"

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

	seconds, err := strconv.Atoi(os.Getenv("REFRESH_TIME_SECONDS"))
	if err != nil {
		log.Fatal("error invalid REFRESH_TIME_SECONDS")
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

	service := NewService(
		solarEdgeClient,
		relay,
		time.Duration(seconds)*time.Second,
	)

	err = service.Run()

	log.Fatal("error: ", err.Error())
}
