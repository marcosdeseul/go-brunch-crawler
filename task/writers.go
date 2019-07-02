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

func urlWriterWithSubscribeNo(url t.URL, subscribeNo int) t.URL {
	return t.URL(fmt.Sprintf("%s&subscribeNo=%d", string(url), subscribeNo))
}

func fetchWriters(url t.URL) ([]Writer, error) {
	first, _ := http.GetData(url)
	var data DataWriter
	marshalled, _ := json.Marshal(first)
	json.Unmarshal(marshalled, &data)
	results := []Writer{}
	results = append(results, data.List...)
	end := data.MoreList
	for end {
		last := results[len(results)-1]
		subscribeNo := last.SubscribeNo
		fmt.Printf("subscribeNo: %d\n", subscribeNo)
		next, _ := http.GetData(urlWriterWithSubscribeNo(url, subscribeNo))
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

func checkTodayWriter(profileID t.ProfileID) (bool, string) {
	return checkTodayFile(profileID, "writer", "csv")
}

// CrawlWriter fetches data when there is no file, otherwise it creates a file after fetching
func CrawlWriter(profileID t.ProfileID, url t.URL) ([]Writer, error) {
	found, fileName := checkTodayWriter(profileID)
	var writers []Writer
	if found {
		fmt.Printf("Today's [Writer] file is found for [%s]\n", profileID)
		csvFile, _ := os.Open(fileName)
		defer csvFile.Close()
		gocsv.Unmarshal(csvFile, &writers)
	} else {
		fmt.Printf("There is no [Writer] file found for [%s]\n", profileID)
		newpath := filepath.Join(".", "output")
		os.MkdirAll(newpath, os.ModePerm)
		writers, _ = fetchWriters(url)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("ERR: %s\n", err)
		}
		defer file.Close()
		gocsv.Marshal(&writers, file)
	}
	return writers, nil
}

type DataWriter struct {
	List       []Writer `json:"list"`
	TotalCount int      `json:"totalCount"`
	MoreList   bool     `json:"moreList"`
}
type Writer struct {
	UserID        string `json:"userId" csv:"userId"`
	UserName      string `json:"userName" csv:"userName"`
	UserImage     string `json:"userImage" csv:"userImage"`
	ProfileID     string `json:"profileId" csv:"profileId"`
	Description   string `json:"description" csv:"description"`
	ArticleCount  int    `json:"articleCount" csv:"articleCount"`
	WriterCount   int    `json:"writerCount" csv:"writerCount"`
	FollowerCount int    `json:"followerCount" csv:"followerCount"`
	CreateTime    int64  `json:"createTime" csv:"createTime"`
	SubscribeNo   int    `json:"subscribeNo" csv:"subscribeNo"`
	MyFollower    bool   `json:"myFollower" csv:"myFollower"`
	MyWriter      bool   `json:"myWriter" csv:"myWriter"`
}
