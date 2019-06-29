package task

import (
	"encoding/json"
	"fmt"

	"github.com/marcosdeseul/go-brunch-crawler/lib/http"
	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
)

type Data struct {
	List       []Article `json:"list"`
	TotalCount int       `json:"totalCount"`
	MoreList   bool      `json:"moreList"`
}

type Article struct {
	UserID                t.UserID `json:"userId"`
	Title                 string   `json:"title"`
	No                    int16    `json:"no"`
	LikeCount             int16    `json:"likeCount"`
	CommentCount          int16    `json:"cikeCount"`
	SocialShareTotalCount int16    `json:"socialSahreTotalCount"`
	ReadSeconds           int16    `json:"readSeconds"`
	CreateTime            int64    `json:"createTime"`
	UpdateTime            int64    `json:"updateTime"`
	PublishTime           int64    `json:"publishTime"`
}

func GetArticles(url t.URL) ([]Article, error) {
	first, _ := http.GetData(url)
	var data Data
	marshalled, _ := json.Marshal(first)
	json.Unmarshal(marshalled, &data)
	results := data.List
	end := data.MoreList
	for end {
		last := results[len(results)-1]
		createdAt := last.CreateTime
		next, _ := http.GetData(t.URL(fmt.Sprintf("%s?lastTime=%d", string(url), createdAt)))
		marshalled, _ := json.Marshal(next)
		json.Unmarshal(marshalled, &data)
		results = append(results, data.List...)
		end = data.MoreList
	}
	return results, nil
}
