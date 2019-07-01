package main

import (
	"fmt"

	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
	"github.com/marcosdeseul/go-brunch-crawler/lib/url"
	"github.com/marcosdeseul/go-brunch-crawler/task"
)

var (
	profileID t.ProfileID
	userID    t.UserID
	listSize  uint8

	urlProfile   t.URL
	urlArticle   t.URL
	urlMagazine  t.URL
	urlWriters   t.URL
	urlFollowers t.URL

	profile task.DataProfile
)

func init() {
	profileID = "imagineer"
	listSize = 100
	urlProfile = url.Profile(profileID)
	urlArticle = url.Article(profileID)
	urlMagazine = url.Magazine(profileID)

	profile, _ = task.CrawlProfile(profileID, urlProfile)
	userID = t.UserID(fmt.Sprintf("%v", profile.UserID))
	urlWriters = url.Writers(userID, 20)
	urlFollowers = url.Followers(userID, 20)
}

func main() {
	task.CrawlArticle(profileID, urlArticle)
}
