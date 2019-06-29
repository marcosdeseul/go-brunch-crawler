package main

import (
	"fmt"

	"github.com/marcosdeseul/go-brunch-crawler/lib/http"
	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
	"github.com/marcosdeseul/go-brunch-crawler/lib/url"
	u "github.com/marcosdeseul/go-brunch-crawler/lib/util"
)

var (
	profileID t.ProfileID
	userID    t.UserID
	listSize  uint8

	urlProfile      t.URL
	urlArticle      t.URL
	urlMagazine     t.URL
	urlSubscription t.URL // base of writers, followers
	urlWriters      t.URL
	urlFollowers    t.URL

	profile map[string]interface{}
)

func init() {
	profileID = "imagineer"
	listSize = 100
	urlProfile = url.Profile(profileID)
	urlArticle = url.Article(profileID)
	urlMagazine = url.Magazine(profileID)
	urlSubscription = url.Subscription()

	profileData, _ := http.GetData(urlProfile)
	profile = profileData
	userID = t.UserID(fmt.Sprintf("%v", profile["userId"]))
	urlWriters = url.Writers(userID, 20)
	urlFollowers = url.Followers(userID, 20)
}

func main() {
	article, _ := http.GetData(urlArticle)
	magazine, _ := http.GetData(urlMagazine)
	writers, _ := http.GetData(urlWriters)
	followers, _ := http.GetData(urlFollowers)
	u.PrettyPrint(article)
	u.PrettyPrint(magazine)
	u.PrettyPrint(writers)
	u.PrettyPrint(followers)
}
