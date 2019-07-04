package main

import (
	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
	w "github.com/marcosdeseul/go-brunch-crawler/task"
)

var profileID t.ProfileID
var fields []w.UserField

func init() {
	profileID = "imagineer"
	fields = []w.UserField{
		w.FieldArticle,
		w.FieldMagazine,
		w.FieldWriter,
		w.FieldFollower,
	}
}

func main() {
	w.GetUser(profileID, fields)
}
