package main

import (
	"fmt"

	"github.com/marcosdeseul/go-brunch-crawler/lib/http"
	"github.com/marcosdeseul/go-brunch-crawler/lib/util"
)

var (
	profileID       string
	userID          string
	listSize        uint8
	urlProfile      string
	urlArticle      string
	urlMagazine     string
	urlSubscription string // base of writers, followers
	urlWriters      string
	urlFollowers    string
)

func init() {
	profileID = "imagineer"
	listSize = 100
	v1 := "v1"
	v2 := "v2"
	v3 := "v3"
	domain := "https://api.brunch.co.kr/"
	urlProfile = fmt.Sprintf("%s/%s/profile/@%s", domain, v1, profileID)
	urlArticle = fmt.Sprintf("%s/%s/article/@%s", domain, v2, profileID)
	urlMagazine = fmt.Sprintf("%s/%s/magazine/@%s", domain, v3, profileID)
	urlSubscription = fmt.Sprintf("%s/%s/subscription/", domain, v2)

	profile := http.GetData(urlProfile)
	userID = fmt.Sprintf("%v", profile["userId"])
	urlWriters = fmt.Sprintf("%s"+"user/@@%s/writers?listSize=%d", urlSubscription, userID, listSize)
	urlFollowers = fmt.Sprintf("%s"+"user/@@%s/followers?listSize=%d", urlSubscription, userID, listSize)
}

func main() {
	article := http.GetData(urlArticle)
	magazine := http.GetData(urlMagazine)
	writers := http.GetData(urlWriters)
	followers := http.GetData(urlFollowers)
	util.PrettyPrint(article)
	util.PrettyPrint(magazine)
	util.PrettyPrint(writers)
	util.PrettyPrint(followers)
}
