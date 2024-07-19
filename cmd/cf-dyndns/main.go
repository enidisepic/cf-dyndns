// Main package of cf-dyndns. Only houses
// the necessary bootstrapping functionality
// for cf-dyndns to work.
package main

import (
	"github.com/enidisepic/cf-dyndns/internal/helpers/anysrc"
	"github.com/enidisepic/cf-dyndns/internal/helpers/cloudflare"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

func run() {
	ipAddress, err := anysrc.GetCurrentIPAddress()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Got IP:", ipAddress)

	err = cloudflare.UpdateEntry(ipAddress)
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
