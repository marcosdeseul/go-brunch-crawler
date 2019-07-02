package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
)

func GetData(url t.URL) (map[string]interface{}, error) {
	return getDataWithRetry(string(url), 5)
}

func getDataWithRetry(url string, retry uint8) (map[string]interface{}, error) {
	body, err := getBody(url, retry)
	return extractData(body), err
}

func getBody(url string, retry uint8) ([]byte, error) {
	if retry < 0 {
		return nil, fmt.Errorf("Out of retries")
	}
	resp, err := http.Get(string(url))
	if err != nil {
		time.Sleep(1 * time.Millisecond)
		return getBody(url, retry-1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		time.Sleep(1 * time.Millisecond)
		return getBody(url, retry-1)
	}
	return body, nil
}

func getDataWithLastTime(url string, username string, lastTime string, retry uint8) (map[string]interface{}, error) {
	formattedURL := fmt.Sprintf("%s@%s?lastTime=%s", url, username, lastTime)
	body, err := getBody(formattedURL, retry)
	return extractData(body), err
}

func extractData(body []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	data, _ := result["data"].(map[string]interface{})
	return data
}
