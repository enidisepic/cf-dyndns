// Package http houses functions related to making the
// http package easier to use for our use case.
package http

import (
	"encoding/json"
	"errors"
	"github.com/enidisepic/cf-dyndns/internal/util"
	"net/http"
)

// Get makes an HTTP GET request and returns the JSON response
// marshaled as T
func Get[T any](url string) (T, error) {
	response, err := http.Get(url)
	if err != nil {
		return util.Zero[T](), errors.New("error requesting " + url)
	}
	defer closeResponseBody(response.Body)

	unmarshalledResponse, err := unmarshalResponse[T](response)
	if err != nil {
		return util.Zero[T](), err
	}

	return unmarshalledResponse, nil
}

// Patch makes an HTTP POST request. It takes in an any
// T and V, T is what the request response will be
// marshaled into, V is the request body
func Patch[T any, V any](
	url string,
	bearerToken string,
	body T,
) (V, error) {
	marshalledRequest, err := json.Marshal(body)
	if err != nil {
		return util.Zero[V](), errors.New("error marshalling request body")
	}

	response, err := doPatchRequest(url, marshalledRequest, bearerToken)
	if err != nil {
		return util.Zero[V](), err
	}
	defer closeResponseBody(response.Body)

	unmarshalledResponse, err := unmarshalResponse[V](response)
	if err != nil {
		return util.Zero[V](), err
	}

	return unmarshalledResponse, nil
}
