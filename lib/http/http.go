package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetData(url string) map[string]interface{} {
	body := getBody(url)
	return extractData(body)
}

func getBody(url string) []byte {
	resp, err := http.Get(url)
	// TODO: handle error here
	if err == nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func getDataWithLastTime(url string, username string, lastTime string) map[string]interface{} {
	formattedURL := fmt.Sprintf("%s@%s?lastTime=%s", url, username, lastTime)
	body := getBody(formattedURL)
	return extractData(body)
}

func extractData(body []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	data := result["data"].(map[string]interface{})
	return data
}
