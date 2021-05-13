package url

import (
	"net/url"
)

func Valid(link string) bool {
	u, err := url.Parse(link)
	return err == nil && u.Scheme != "" && u.Host != ""
}
