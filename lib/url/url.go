package url

import (
	"fmt"

	"github.com/marcosdeseul/go-brunch-crawler/lib/t"
)

// url related consts
const (
	DOMAIN       string = "https://api.brunch.co.kr"
	V1           string = "v1"
	V2           string = "v2"
	V3           string = "v3"
	PROFILE      string = "profile"
	ARTICLE      string = "article"
	MAGAZINE     string = "magazine"
	SUBSCRIPTION string = "subscription"
)

type urlGenerator interface {
	build() t.URL
}

type urlInfo struct {
	Domain  string
	Version string
	Target  string
	Extra   string
}

func (info urlInfo) build() t.URL {
	return t.URL(fmt.Sprintf("%s/%s/%s%s", info.Domain, info.Version, info.Target, info.Extra))
}

// Profile return a complete url
func Profile(profileID t.ProfileID) t.URL {
	var info urlGenerator = urlInfo{
		Domain:  DOMAIN,
		Version: V1,
		Target:  PROFILE,
		Extra:   fmt.Sprintf("/@%s", profileID),
	}
	return info.build()
}

// Article return a complete url
func Article(profileID t.ProfileID) t.URL {
	var info urlGenerator = urlInfo{
		Domain:  DOMAIN,
		Version: V2,
		Target:  ARTICLE,
		Extra:   fmt.Sprintf("/@%s", profileID),
	}
	return info.build()
}

// Magazine return a complete url
func Magazine(profileID t.ProfileID) t.URL {
	var info urlGenerator = urlInfo{
		Domain:  DOMAIN,
		Version: V3,
		Target:  MAGAZINE,
		Extra:   fmt.Sprintf("/@%s", profileID),
	}
	return info.build()
}

// Writers return a complete url
func Writers(userID t.UserID, listSize int8) t.URL {
	var info urlGenerator = urlInfo{
		Domain:  DOMAIN,
		Version: V2,
		Target:  SUBSCRIPTION,
		Extra:   fmt.Sprintf("/user/@@%s/writers?listSize=%d", userID, listSize),
	}
	return info.build()
}

// Followers return a complete url
func Followers(userID t.UserID, listSize int8) t.URL {
	var info urlGenerator = urlInfo{
		Domain:  DOMAIN,
		Version: V2,
		Target:  SUBSCRIPTION,
		Extra:   fmt.Sprintf("/user/@@%s/followers?listSize=%d", userID, listSize),
	}
	return info.build()
}
