package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

func CheckAdmin(email string) (bool, error) {
	client := resty.New()
	encodeEmail := url.QueryEscape(email)
	url := fmt.Sprintf("http://127.0.0.1:9001/check/%s", encodeEmail)

	log.Printf("sending request to url: %+v", url)

	resp, err := client.R().Get(url)

	if err != nil {
		log.Printf("Error making request: %+v", err)
		return false, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		log.Printf("User not found or API endpoint not accessible: %s", url)
		return false, fmt.Errorf("user not found or API endpoint not accessible")
	} else if resp.StatusCode() != http.StatusOK {
		log.Printf("Unexpected status code: %d, body: %s", resp.StatusCode(), resp.Body())
		return false, fmt.Errorf("unexpected status code: %d, body: %+v", resp.StatusCode(), resp.Body())
	}

	log.Printf("Response Body: %s", resp.Body())

	var result map[string]bool
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return false, fmt.Errorf("failed to unmarshal response: %+v, body: %s", err, resp.Body())
	}

	isAdmin, ok := result["isadmin"]
	if !ok {
		return false, fmt.Errorf("key 'isAdmin' not found in response")
	}

	return isAdmin, nil

}
