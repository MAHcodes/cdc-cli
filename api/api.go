package api

import (
	"fmt"
	"io"
	"net/http"
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
