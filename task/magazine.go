package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/marcosdeseul/go-brunch-crawler/lib/http"
	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
)

func fetchMagazines(url t.URL) ([]MagazineJSON, error) {
	magazineData, _ := http.GetData(url)
	var magazines DataMagazine
	marshalled, _ := json.Marshal(magazineData)
	json.Unmarshal(marshalled, &magazines)
	results := magazines.List
	sort.Slice(results, func(i, j int) bool {
		return results[i].No < results[j].No
	})
	return results, nil
}

func checkTodayMagazine(profileID t.ProfileID) (bool, string) {
	return checkTodayFile(profileID, "magazine", "csv")
}

// CrawlMagazine fetches data when there is no file, otherwise it creates a file after fetching
func CrawlMagazine(profileID t.ProfileID, url t.URL) ([]Magazine, error) {
	found, fileName := checkTodayMagazine(profileID)
	var magazines []Magazine
	if found {
		fmt.Printf("Today's [Magazine] file is found for [%s]\n", profileID)
		csvFile, _ := os.Open(fileName)
		defer csvFile.Close()
		gocsv.Unmarshal(csvFile, &magazines)
	} else {
		fmt.Printf("There is no [Magazine] file found for [%s]\n", profileID)
		newpath := filepath.Join(".", "output")
		os.MkdirAll(newpath, os.ModePerm)
		magazineJSON, _ := fetchMagazines(url)
		magazines = converMagazineJSON(magazineJSON)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("ERR: %s", err)
		}
		defer file.Close()
		gocsv.Marshal(&magazines, file)
	}
	return magazines, nil
}

func converMagazineJSON(magazineJSON []MagazineJSON) []Magazine {
	var magazines []Magazine
	for _, json := range magazineJSON {
		var tagList []string
		for _, tag := range json.TagList {
			tagList = append(tagList, fmt.Sprintf("%d|%s", tag.TagID, tag.Name))
		}
		m := Magazine{
			No:                   json.No,
			Title:                json.Title,
			MagazineAddress:      json.MagazineAddress,
			CreateTime:           json.CreateTime,
			UpdateTime:           json.UpdateTime,
			ShareCountUpdateTime: json.ShareCountUpdateTime,
			ArticleCount:         json.ArticleCount,
			FollowerCount:        json.FollowerCount,
			AllowJoin:            json.AllowJoin,
			BgColor:              json.BgColor,
			Brunchbook:           json.Brunchbook,
			MagazineStatus:       json.MagazineStatus,
			LikeCount:            json.LikeCount,
			TagList:              strings.Join(tagList, ","),
			AuthorCount:          json.AuthorCount,
			ShareCount:           json.ShareCount,
			MagazineBookCover:    json.MagazineBookCover,
		}
		magazines = append(magazines, m)
	}
	return magazines
}

type DataMagazine struct {
	MoreCount  int            `json:"moreCount"`
	List       []MagazineJSON `json:"list"`
	TotalCount int            `json:"totalCount"`
	MoreList   bool           `json:"moreList"`
}
type TagList struct {
	TagID int    `json:"tagId" csv:"tagId"`
	Name  string `json:"name" csv:"name"`
	Cnt   int    `json:"cnt" csv:"cnt"`
}
type MagazineJSON struct {
	No                   int       `json:"no"`
	Title                string    `json:"title"`
	MagazineAddress      string    `json:"magazineAddress"`
	CreateTime           int64     `json:"createTime"`
	UpdateTime           int64     `json:"updateTime"`
	ShareCountUpdateTime int64     `json:"shareCountUpdateTime"`
	ArticleCount         int       `json:"articleCount"`
	FollowerCount        int       `json:"followerCount"`
	AllowJoin            bool      `json:"allowJoin"`
	BgColor              string    `json:"bgColor"`
	Brunchbook           bool      `json:"brunchbook"`
	MagazineStatus       string    `json:"magazineStatus"`
	LikeCount            int       `json:"likeCount"`
	TagList              []TagList `json:"tagList"`
	AuthorCount          int       `json:"authorCount"`
	ShareCount           int       `json:"shareCount"`
	MagazineBookCover    bool      `json:"magazineBookCover"`
}
type Magazine struct {
	No                   int    `csv:"no"`
	Title                string `csv:"title"`
	MagazineAddress      string `csv:"magazineAddress"`
	CreateTime           int64  `csv:"createTime"`
	UpdateTime           int64  `csv:"updateTime"`
	ShareCountUpdateTime int64  `csv:"shareCountUpdateTime"`
	ArticleCount         int    `csv:"articleCount"`
	FollowerCount        int    `csv:"followerCount"`
	AllowJoin            bool   `csv:"allowJoin"`
	BgColor              string `csv:"bgColor"`
	Brunchbook           bool   `csv:"brunchbook"`
	MagazineStatus       string `csv:"magazineStatus"`
	LikeCount            int    `csv:"likeCount"`
	TagList              string `csv:"tagList"`
	AuthorCount          int    `csv:"authorCount"`
	ShareCount           int    `csv:"shareCount"`
	MagazineBookCover    bool   `csv:"magazineBookCover"`
}
