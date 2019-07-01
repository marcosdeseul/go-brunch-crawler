package task

import (
	"github.com/marcosdeseul/go-brunch-crawler/lib/file"
	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
)

func checkTodayFile(profileID t.ProfileID, category string, extension string) (bool, string) {
	fileName := file.GetFileName("output", profileID, category, extension)
	return file.FindFile(fileName), fileName
}
