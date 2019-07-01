package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/marcosdeseul/go-brunch-crawler/lib/http"
	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
)

func checkTodayProfile(profileID t.ProfileID) (bool, string) {
	return checkTodayFile(profileID, "profile", "json")
}

func CrawlProfile(profileID t.ProfileID, url t.URL) (DataProfile, error) {
	found, fileName := checkTodayProfile(profileID)
	var profile DataProfile
	if found {
		fmt.Printf("Today's [Profile] file is found for [%s]\n", profileID)
		file, _ := ioutil.ReadFile(fileName)
		json.Unmarshal(file, &profile)
	} else {
		fmt.Printf("There is no [Profile] file found for [%s]\n", profileID)
		newpath := filepath.Join(".", "output")
		os.MkdirAll(newpath, os.ModePerm)
		profileData, _ := http.GetData(url)
		marshalled, _ := json.Marshal(profileData)
		json.Unmarshal(marshalled, &profile)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("ERR: %s", err)
		}
		defer file.Close()
		file.Write(marshalled)
	}
	return profile, nil
}

type DataProfile struct {
	UserID                string                `json:"userId"`
	UserName              string                `json:"userName"`
	ProfileID             string                `json:"profileId"`
	Description           string                `json:"description"`
	CreateTime            int64                 `json:"createTime"`
	Status                string                `json:"status"`
	ArticleCount          int                   `json:"articleCount"`
	SocialUser            bool                  `json:"socialUser"`
	WriterCount           int                   `json:"writerCount"`
	FollowerCount         int                   `json:"followerCount"`
	MagazineCount         int                   `json:"magazineCount"`
	Company               string                `json:"company"`
	DescriptionDetail     string                `json:"descriptionDetail"`
	ProfileNotiExpireTime int64                 `json:"profileNotiExpireTime"`
	UserSns               UserSns               `json:"userSns"`
	BookStoreBookList     []BookStoreBookList   `json:"bookStoreBookList"`
	ProfileCategoryList   []ProfileCategoryList `json:"profileCategoryList"`
	AcceptPropose         bool                  `json:"acceptPropose"`
	BrunchActivityList    []BrunchActivityList  `json:"brunchActivityList"`
	EmailValidation       bool                  `json:"emailValidation"`
	Author                bool                  `json:"author"`
}
type UserSns struct {
	UserID     string      `json:"userId"`
	Website    string      `json:"website"`
	Facebook   string      `json:"facebook"`
	Twitter    interface{} `json:"twitter"`
	Instagram  string      `json:"instagram"`
	CreateTime int64       `json:"createTime"`
	UpdateTime int64       `json:"updateTime"`
}
type BookStoreBookList struct {
	No            int    `json:"no"`
	Section       string `json:"section"`
	Title         string `json:"title"`
	UserID        string `json:"userId"`
	PcLink        string `json:"pcLink"`
	MobileLink    string `json:"mobileLink"`
	PublisherName string `json:"publisherName"`
	PublisherLink string `json:"publisherLink"`
	Description   string `json:"description"`
	Sentence      string `json:"sentence"`
	PublishTime   int64  `json:"publishTime"`
	SaveTime      int64  `json:"saveTime"`
	MagazineNo    int    `json:"magazineNo"`
	WriterName    string `json:"writerName"`
	DaumBookID    string `json:"daumBookId"`
}
type KeywordList struct {
	No        int    `json:"no"`
	Keyword   string `json:"keyword"`
	Sequence  int    `json:"sequence"`
	KeywordNo int    `json:"keywordNo"`
}
type ProfileCategoryList struct {
	Category     string        `json:"category"`
	CategoryNo   int           `json:"categoryNo"`
	CategoryName string        `json:"categoryName"`
	KeywordList  []KeywordList `json:"keywordList"`
}
type BrunchActivityList struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	URL      string `json:"url"`
}
