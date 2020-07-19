package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	unicode2 "golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"pujichun.com/learnspider/MaoYanPiaoFang/parser"
)

func GetUid() (uid string) {
	client := http.DefaultClient
	request, err := http.NewRequest("GET", "https://piaofang.maoyan.com/rankings/year?year=2020&limit=100&tab=1", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	request.Header.Add("Host", "piaofang.maoyan.com")
	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		panic(err)
	}
	body := bufio.NewReader(response.Body)
	e := determineCharset(body)
	newBody := transform.NewReader(body, e.NewDecoder())
	content, err := ioutil.ReadAll(newBody)
	if err != nil {
		panic(err)
	}
	return parser.Uid(content)
}

func Fetch(u, uid string){
	client := http.DefaultClient
	request, err := http.NewRequest("GET", u, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	request.Header.Add("Host", "piaofang.maoyan.com")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Referer", "https://piaofang.maoyan.com/rankings/year?year=2020&limit=100&tab=1")
	request.Header.Add("Uid", uid)
	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	//if err != nil {
	//	return nil, err
	//}
	defer response.Body.Close()
	//if response.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("request %s status code is %d", u, response.StatusCode)
	//}
	bodyReader := bufio.NewReader(response.Body)
	e := determineCharset(bodyReader)
	encodeBody := transform.NewReader(bodyReader, e.NewDecoder())
	//return ioutil.ReadAll(encodeBody)
	content, err := ioutil.ReadAll(encodeBody)
	rank := parser.Rank(content)
	fmt.Println(rank)

}

func determineCharset(i io.Reader) encoding.Encoding {
	resp, err := bufio.NewReader(i).Peek(1024)
	if err != nil {
		return unicode2.UTF8
	}
	e, _, _ := charset.DetermineEncoding(resp, "")
	return e
}
