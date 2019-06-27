package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	username := "imagineer"
	v1 := "v1"
	// v2 := "v2"
	// v3 := "v3"
	domain := "https://api.brunch.co.kr/"
	urlProfile := fmt.Sprintf("%s/%s/profile/", domain, v1)
	// url_article := fmt.Sprintf("%s/%s/article/", domain, v2)
	// https://api.brunch.co.kr/v2/article/@imagineer?lastTime=1525179400000

	profile := getData(urlProfile, username)
	// article := getData(url_article, username)

	prettyPrint(profile)
	// prettyPrint(article)
}

func prettyPrint(data map[string]interface{}) {
	prettier, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(prettier))
}

func getBody(url string) []byte {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func getData(url string, username string) map[string]interface{} {
	formattedURL := fmt.Sprintf("%s@%s", url, username)
	body := getBody(formattedURL)
	return extractData(body)
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
