package file

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
)

func FindFile(name string) bool {
	var result bool
	_ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if path == name {
			result = true
		}
		return nil
	})
	return result
}

func GetFileName(folder string, profileID t.ProfileID, category string, ext string) string {
	t := time.Now()
	return fmt.Sprintf("%s/%s-%s-%s.%s", folder, profileID, t.Format("20060102"), category, ext)
}
