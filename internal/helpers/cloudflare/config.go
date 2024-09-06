package cloudflare

import (
	"errors"
	"fmt"
	"github.com/enidisepic/cf-dyndns/internal/util"
	"os"
)

type entryUpdateRequestConfig struct {
	APIKEY    string
	APIURL    string
	ZoneID    string
	EntryName string
	Proxied   bool
}

func getConfig() (entryUpdateRequestConfig, error) {
	apiKey := os.Getenv("CF_API_KEY")
	if apiKey == "" {
		return util.Zero[entryUpdateRequestConfig](), errors.New("CF_API_KEY environment variable not set")
	}
	zoneID := os.Getenv("CF_ZONE_ID")
	if zoneID == "" {
		return util.Zero[entryUpdateRequestConfig](), errors.New("CF_ZONE_ID environment variable not set")
	}
	entryID := os.Getenv("CF_ENTRY_ID")
	if entryID == "" {
		return util.Zero[entryUpdateRequestConfig](), errors.New("CF_ENTRY_ID environment variable not set")
	}
	entryName := os.Getenv("CF_ENTRY_NAME")
	if entryName == "" {
		return util.Zero[entryUpdateRequestConfig](), errors.New("CF_ENTRY_NAME environment variable not set")
	}
	
	proxied := os.Getenv("CF_PROXIED")

	cloudflareAPIURL := fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s",
		zoneID,
		entryID,
	)

	return entryUpdateRequestConfig{
		APIKEY:    apiKey,
		APIURL:    cloudflareAPIURL,
		ZoneID:    zoneID,
		EntryName: entryName,	
		Proxied:   proxied != "",
	}, nil
}
