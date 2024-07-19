// Package anysrc houses the wrapper for
// the anysrc.net API.
package anysrc

import (
	"github.com/enidisepic/cf-dyndns/internal/helpers/http"
	"log"
)

type response struct {
	IPAddress string `json:"clientip"`
}

// GetCurrentIPAddress fetches the current public IP
// and returns it.
func GetCurrentIPAddress() (string, error) {
	apiURL := "https://ip4.anysrc.net/json"

	log.Println("Getting IP")

	apiResponse, err := http.Get[response](apiURL)
	if err != nil {
		return "", err
	}

	return apiResponse.IPAddress, nil
}
