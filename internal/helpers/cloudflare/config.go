package cloudflare

import (
	"errors"
	"fmt"
	"github.com/enidisepic/cf-dyndns/internal/util"
	"os"
)

type entryUpdateRequestConfig struct {
	ApiKey    string
	ApiUrl    string
	ZoneId    string
	EntryName string
}

func getConfig() (entryUpdateRequestConfig, error) {
	apiKey := os.Getenv("CF_API_KEY")
	if apiKey == "" {
		return util.Zero[entryUpdateRequestConfig](), errors.New("CF_API_KEY environment variable not set")
	}
	zoneId := os.Getenv("CF_ZONE_ID")
	if zoneId == "" {
		return util.Zero[entryUpdateRequestConfig](), errors.New("CF_ZONE_ID environment variable not set")
	}
	entryId := os.Getenv("CF_ENTRY_ID")
	if entryId == "" {
		return util.Zero[entryUpdateRequestConfig](), errors.New("CF_ENTRY_ID environment variable not set")
	}
	entryName := os.Getenv("CF_ENTRY_NAME")
	if entryName == "" {
		return util.Zero[entryUpdateRequestConfig](), errors.New("CF_ENTRY_NAME environment variable not set")
	}

	cloudflareApiUrl := fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s",
		zoneId,
		entryId,
	)

	return entryUpdateRequestConfig{
		ApiKey:    apiKey,
		ApiUrl:    cloudflareApiUrl,
		ZoneId:    zoneId,
		EntryName: entryName,
	}, nil
}
