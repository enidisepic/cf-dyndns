package http_helper

import (
	"encoding/json"
	"errors"
	"github.com/enidisepic/cf-dyndns/internal/util"
	"net/http"
)

func Get[T interface{}](url string) (T, error) {
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

func Patch[T interface{}, V interface{}](
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
