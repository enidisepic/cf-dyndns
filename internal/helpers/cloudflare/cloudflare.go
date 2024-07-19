// Package cloudflare houses helper functions
// for the Cloudflare API.
package cloudflare

import (
	"errors"
	"github.com/enidisepic/cf-dyndns/internal/helpers/http"
	"log"
)

func UpdateEntry(ipAddress string) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	cloudflareEntryCreateRequest := createEntryUpdateRequest(
		config,
		ipAddress,
	)

	log.Println("Updating DNS entry for:", config.EntryName)
	cloudflareEntryCreateResponse, err := http.Patch[
		entryUpdateRequest,
		entryUpdateResponse,
	](
		config.ApiUrl,
		config.ApiKey,
		cloudflareEntryCreateRequest,
	)
	if err != nil {
		return err
	} else if !cloudflareEntryCreateResponse.Success {
		return errors.New("error while updating dns entry")
	}

	return nil
}
