package task

import (
	"fmt"

	"github.com/marcosdeseul/go-brunch-crawler/lib/file"
	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
	"github.com/marcosdeseul/go-brunch-crawler/lib/url"
)

func checkTodayFile(profileID t.ProfileID, category string, extension string) (bool, string) {
	fileName := file.GetFileName("output", profileID, category, extension)
	return file.FindFile(fileName), fileName
}

type UserField string

const FieldArticle UserField = "Article"
const FieldMagazine UserField = "Magazine"
const FieldFollower UserField = "Follower"
const FieldWriter UserField = "Writer"

// User contains all data about a user
type User struct {
	Fields []UserField

	Profile   // default
	Articles  []Article
	Magazines []Magazine
	Followers []Follower
	Writers   []Writer
}

func GetUser(profileID t.ProfileID, fields []UserField) (User, []error) {
	var errors []error

	urlProfile := url.Profile(profileID)
	profile, _ := CrawlProfile(profileID, urlProfile)
	userID := t.UserID(fmt.Sprintf("%v", profile.UserID))
	listSize := uint(100)

	user := User{
		Fields:  fields,
		Profile: profile,
	}
	for _, f := range fields {
		switch f {
		case FieldArticle:
			urlArticle := url.Article(profileID)
			user.Articles, _ = CrawlArticle(profileID, urlArticle)
		case FieldMagazine:
			urlMagazine := url.Magazine(profileID)
			user.Magazines, _ = CrawlMagazine(profileID, urlMagazine)
		case FieldFollower:
			urlFollower := url.Followers(userID, listSize)
			user.Followers, _ = CrawlFollower(profileID, urlFollower)
		case FieldWriter:
			urlWriter := url.Writers(userID, listSize)
			user.Writers, _ = CrawlWriter(profileID, urlWriter)
		default:
			fmt.Printf("Field [%s] is not handled", f)
		}
	}
	return user, errors
}
