package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetJSON[T any](url string) ([]T, error) {
	client := &http.Client{
		Timeout: 5*time.Second,
	}
	
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var data []T
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
