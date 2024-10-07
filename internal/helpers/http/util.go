package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/enidisepic/cf-dyndns/internal/util"
	"io"
	"log"
	"net/http"
)

func closeResponseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Println("error closing response body")
	}
}

func doPatchRequest(url string, body []byte, bearerToken string) (*http.Response, error) {
	request, err := http.NewRequest(
		"PATCH",
		url,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New("error creating request")
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+bearerToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.New("error sending patch request")
	}

	return response, nil
}

func unmarshalResponse[T any](response *http.Response) (T, error) {
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return util.Zero[T](), errors.New("error reading response body")
	}

	var unmarshalledResponse T
	if json.Unmarshal(responseBody, &unmarshalledResponse) != nil {
		return util.Zero[T](), errors.New("error unmarshalling response body")
	}

	return unmarshalledResponse, nil
}
