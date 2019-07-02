package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/gocarina/gocsv"
	"github.com/marcosdeseul/go-brunch-crawler/lib/http"
	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
)

func urlFollowerWithSubscribeNo(url t.URL, subscribeNo int) t.URL {
	return t.URL(fmt.Sprintf("%s&subscribeNo=%d", string(url), subscribeNo))
}

func fetchFollowers(url t.URL) ([]Follower, error) {
	first, _ := http.GetData(url)
	var data DataFollower
	marshalled, _ := json.Marshal(first)
	json.Unmarshal(marshalled, &data)
	results := []Follower{}
	results = append(results, data.List...)
	end := data.MoreList
	for end {
		last := results[len(results)-1]
		subscribeNo := last.SubscribeNo
		fmt.Printf("subscribeNo: %d\n", subscribeNo)
		next, _ := http.GetData(urlFollowerWithSubscribeNo(url, subscribeNo))
		marshalled, _ := json.Marshal(next)
		json.Unmarshal(marshalled, &data)
		results = append(results, data.List...)
		end = data.MoreList
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].CreateTime > results[j].CreateTime
	})
	return results, nil
}

func checkTodayFollower(profileID t.ProfileID) (bool, string) {
	return checkTodayFile(profileID, "follower", "csv")
}

// CrawlFollower fetches data when there is no file, otherwise it creates a file after fetching
func CrawlFollower(profileID t.ProfileID, url t.URL) ([]Follower, error) {
	found, fileName := checkTodayFollower(profileID)
	var followers []Follower
	if found {
		fmt.Printf("Today's [Follower] file is found for [%s]\n", profileID)
		csvFile, _ := os.Open(fileName)
		defer csvFile.Close()
		gocsv.Unmarshal(csvFile, &followers)
	} else {
		fmt.Printf("There is no [Follower] file found for [%s]\n", profileID)
		newpath := filepath.Join(".", "output")
		os.MkdirAll(newpath, os.ModePerm)
		followers, _ = fetchFollowers(url)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("ERR: %s\n", err)
		}
		defer file.Close()
		gocsv.Marshal(&followers, file)
	}
	return followers, nil
}

type DataFollower struct {
	List       []Follower `json:"list"`
	TotalCount int        `json:"totalCount"`
	MoreList   bool       `json:"moreList"`
}
type Follower struct {
	SubscribeNo   int    `json:"subscribeNo" csv:"subscribeNo"`
	UserID        string `json:"userId" csv:"userId"`
	UserName      string `json:"userName" csv:"userName"`
	ProfileID     string `json:"profileId" csv:"profileId"`
	ArticleCount  int    `json:"articleCount" csv:"articleCount"`
	WriterCount   int    `json:"writerCount" csv:"writerCount"`
	FollowerCount int    `json:"followerCount" csv:"followerCount"`
	CreateTime    int64  `json:"createTime" csv:"createTime"`
	MyFollower    bool   `json:"myFollower" csv:"myFollower"`
	MyWriter      bool   `json:"myWriter" csv:"myWriter"`
	// Description   string `json:"description" csv:"description"`
}
