package anysrc_wrapper

import (
	"github.com/enidisepic/cf-dyndns/internal/http_helper"
	"log"
)

type response struct {
	IpAddress string `json:"clientip"`
}

func GetCurrentIpAddress() (string, error) {
	apiUrl := "https://ip4.anysrc.net/json"

	log.Println("Getting IP")

	apiResponse, err := http_helper.Get[response](apiUrl)
	if err != nil {
		return "", err
	}

	return apiResponse.IpAddress, nil
}
