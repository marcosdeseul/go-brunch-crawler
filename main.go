package main

import (
	"fmt"
)

func main() {
	username := "imagineer"
	v1 := "v1"
	// v2 := "v2"
	// v3 := "v3"
	domain := "https://api.brunch.co.kr/"
	urlProfile := fmt.Sprintf("%s/%s/profile/", domain, v1)
	// urlArticle := fmt.Sprintf("%s/%s/article/", domain, v2)

	profile := getData(urlProfile, username)
	// article := getData(urlArticle, username)

	prettyPrint(profile)
	// prettyPrint(article)
}
