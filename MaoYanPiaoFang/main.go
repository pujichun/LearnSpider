package main

import (
	"pujichun.com/learnspider/MaoYanPiaoFang/fetcher"
)

func main() {
	//str := "{\"totalBox\":\"<span id=\\\"year-box\\\">2020年票房排行榜</span>\\n<span id=\\\"update-time\\\">(截至7月16日 总票房22亿元)</span>\\n\",\"yearList\":\"<ul class=\\\"row\\\" data-com=\\\"hrefTo,href:'/movie/1279731'\\\">\\n    <li class=\\\"col0\\\">1</li>\\n    <li class=\\\"col1\\\">\\n        <p class=\\\"first-line\\\">宠爱</p>\\n        <p class=\\\"second-line\\\">2019-12-31 上映</p>\\n    </li>\\n    <li class=\\\"col2 tr\\\">51006</li>\\n"
	//yearInfo := regexp.MustCompile(`<span id=\\"year-box\\">(\d+)年票房排行榜</span>\\n<span id=\\"update-time\\">\(.*?房(.*?)亿元\)</span>\\n`)
	//match := yearInfo.FindAllSubmatch([]byte(str), -1)
	//fmt.Println(string(match[0][1]), string(match[0][2]))
	urls := fetcher.ParseUrl()
	uid := fetcher.GetUid()
	for _, v := range urls{
		//time.Sleep(time.Second)
		fetcher.Fetch(v, uid)
	}

}
