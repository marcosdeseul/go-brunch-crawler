package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/marcosdeseul/go-brunch-crawler/lib/http"
	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
)

// DataArticle includes a list of articles and meta data
type DataArticle struct {
	List       []Article `json:"list"`
	TotalCount int       `json:"totalCount"`
	MoreList   bool      `json:"moreList"`
}

// Article is a struct to describe article data
type Article struct {
	No                    int16  `json:"no" csv:"no"`
	Title                 string `json:"title" csv:"title"`
	LikeCount             int16  `json:"likeCount" csv:"like"`
	CommentCount          int16  `json:"commentCount" csv:"comment"`
	SocialShareTotalCount int16  `json:"socialShareTotalCount" csv:"socialShareTotal"`
	ReadSeconds           int16  `json:"readSeconds" csv:"readSeconds"`
	CreateTime            int64  `json:"createTime" csv:"createTime"`
	UpdateTime            int64  `json:"updateTime" csv:"updateTime"`
	PublishTime           int64  `json:"publishTime" csv:"publishTime"`
}

func urlArticleWithLastTime(url t.URL, time int64) t.URL {
	return t.URL(fmt.Sprintf("%s?lastTime=%d", string(url), time))
}

func fetchArticles(url t.URL) ([]Article, error) {
	now := int64(time.Now().Unix() * 1000)
	first, _ := http.GetData(urlArticleWithLastTime(url, now))
	var data DataArticle
	marshalled, _ := json.Marshal(first)
	json.Unmarshal(marshalled, &data)
	results := []Article{}
	results = append(results, data.List...)
	end := data.MoreList
	for end {
		last := results[len(results)-1]
		publishedAt := last.PublishTime
		next, _ := http.GetData(urlArticleWithLastTime(url, publishedAt))
		marshalled, _ := json.Marshal(next)
		json.Unmarshal(marshalled, &data)
		results = append(results, data.List...)
		end = data.MoreList
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].No < results[j].No
	})
	return results, nil
}

func checkTodayArticle(profileID t.ProfileID) (bool, string) {
	return checkTodayFile(profileID, "article", "csv")
}

// CrawlArticle fetches data when there is no file, otherwise it creates a file after fetching
func CrawlArticle(profileID t.ProfileID, url t.URL) ([]Article, error) {
	found, fileName := checkTodayArticle(profileID)
	var articles []Article
	if found {
		fmt.Printf("Today's Article file is found for [%s]\n", profileID)
		csvFile, _ := os.Open(fileName)
		defer csvFile.Close()
		gocsv.Unmarshal(csvFile, &articles)
	} else {
		fmt.Printf("There is no Article file found for [%s]\n", profileID)
		newpath := filepath.Join(".", "output")
		os.MkdirAll(newpath, os.ModePerm)
		articles, _ = fetchArticles(url)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("ERR: %s", err)
		}
		defer file.Close()
		gocsv.Marshal(&articles, file)
	}
	return articles, nil
}
