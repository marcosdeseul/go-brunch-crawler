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
	url_profile := fmt.Sprintf("%s/%s/profile/", domain, v1)
	data := getUserProfile(url_profile, username)

	prettyPrint(data)
}

func prettyPrint(data map[string]interface{}) {
	prettier, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(prettier))
}

func getUserProfile(url_profile string, username string) map[string]interface{} {
	url_profile_user := fmt.Sprintf("%s@%s", url_profile, username)
	resp, _ := http.Get(url_profile_user)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return extractData(body)
}

func extractData(body []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	data := result["data"].(map[string]interface{})
	return data
}
