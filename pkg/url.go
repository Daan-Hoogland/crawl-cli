package pkg

import (
	"net/url"
	"strconv"
)

//GenerateURL generates the url from an ip string, port and optional extensions
func GenerateURL(ip string, port int, parameters map[string]string, ssl bool, extension ...string) string {
	urlString := ip + ":" + strconv.Itoa(port)

	if ssl {
		urlString = "https://" + urlString
	} else {
		urlString = "http://" + urlString
	}

	if extension != nil {
		for _, s := range extension {
			urlString = urlString + "/" + s
		}
	}

	if parameters != nil {
		params := url.Values{}
		for key, val := range parameters {
			params.Add(key, val)
		}
		urlString = urlString + params.Encode()
	}

	return urlString
}
