package fetcher

import (
	"fmt"
	"net/url"
	"strconv"
)

func ParseUrl() []string {
	urls := make([]string, 0)
	for i := 1; i <= 10; i++ {
		v := url.Values{}
		v.Add("year", strconv.Itoa(2020-i+1))
		v.Add("limit", "100")
		v.Add("tab", strconv.Itoa(i))
		body := v.Encode()
		result := fmt.Sprint("https://piaofang.maoyan.com/rankings/year?", body)
		urls = append(urls, result)
	}
	return urls
}
