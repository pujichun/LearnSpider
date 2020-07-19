package parser

import (
	"pujichun.com/learnspider/MaoYanPiaoFang/class"
	"regexp"
)

var (
	csrf  = regexp.MustCompile(`<meta name="csrf" content="(.*?)" />`)
	movieRank = regexp.MustCompile(`<ul class=\\"row\\" data-com=\\"hrefTo,href:'(.*?)'\\">\\n\s+<li class=\\"col0\\">(\d+)</li>\\n\s+<li class=\\"col1\\">\\n\s+<p class=\\"first-line\\">(.*?)</p>\\n\s+<p class=\\"second-line\\">(.*?) 上映</p>\\n\s+</li>\\n\s+<li class=\\"col2 tr\\">(\d+)</li>\\n\s+<li class=\\"col3 tr\\">(.*?)</li>\\n\s+<li class=\\"col4 tr\\">(.*?)</li>\\n</ul`)
	yearInfo = regexp.MustCompile(`<span id=\\"year-box\\">(\d+)年票房排行榜</span>\\n<span id=\\"update-time\\">\(.*?房(.*?)亿元\)</span>\\n`)
)

func Uid(content []byte) (uid string) {
	match := csrf.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}


func Rank(content []byte) class.MaoYanRank {
	rank := class.MaoYanRank{}
	match := yearInfo.FindAllSubmatch(content, -1)
	rank.Year = string(match[0][1])
	rank.SumBoxOfficeIncome = string(match[0][2])
	match = movieRank.FindAllSubmatch(content, -1)
	info := class.MovieInfo{}
	for _, m := range match{
		info.Url = string(m[1])
		info.Rank = string(m[2])
		info.MovieName = string(m[3])
		info.BoxOfficeIncome = string(m[4])
		info.AvgEachTicket = string(m[5])
		info.PerformancePeople = string(m[6])
		rank.MovieRank = append(rank.MovieRank, info)
	}
	return rank
}
