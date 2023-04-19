package api

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func FetchEndpoint(endpoint string) ([]byte, error) {
	res, err := http.Get(endpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %v", endpoint, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch URL %s: got status code %d", endpoint, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

func FetchJSON[T interface{}](endpoint string, t *T) (*T, error) {
	j, err := FetchEndpoint(endpoint)
	if err != nil {
		return t, fmt.Errorf("failed to fetch endpoint %s", err)
	}
	if err := json.Unmarshal(j, &t); err != nil {
		return t, fmt.Errorf("failed to unmarshal JSON: %v\n", err)
	}
  return t, nil
}
