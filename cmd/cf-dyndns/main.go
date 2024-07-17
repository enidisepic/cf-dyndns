package main

import (
	"github.com/enidisepic/cf-dyndns/internal/anysrc_wrapper"
	"github.com/enidisepic/cf-dyndns/internal/cloudflare_wrapper"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

func run() {
	ipAddress, err := anysrc_wrapper.GetCurrentIpAddress()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Got IP:", ipAddress)

	err = cloudflare_wrapper.UpdateEntry(ipAddress)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Successfully updated DNS entry")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on global env vars")
	}

	cronSchedule := os.Getenv("CRON_SCHEDULE")
	if cronSchedule == "" {
		log.Println("Cron schedule not set, using default (every 5 minutes)")
		cronSchedule = "@every 5m"
	}
	log.Println("Using cron schedule:", cronSchedule)

	run()
	cronRunner := cron.New()
	_, err = cronRunner.AddFunc(cronSchedule, run)
	if err != nil {
		log.Fatal(err)
		return
	}
	cronRunner.Start()
	select {}
}
