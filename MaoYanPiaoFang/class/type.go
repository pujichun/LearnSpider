package class


type MaoYanRank struct {
	Year string
	SumBoxOfficeIncome string
	MovieRank []MovieInfo
}

type MovieInfo struct {
	Url               string
	Rank              string
	MovieName         string
	StartYear         string
	BoxOfficeIncome   string
	AvgEachTicket     string
	PerformancePeople string
}