package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

func getUrl(url string) (content string, status int) {
	resp, err := http.Get(url)
	if err != nil {
		status = 100
		fmt.Printf("can not open url: %s\nerr: %v", url, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		status = -200
		fmt.Printf("can not get url body.err: %v\n", err)
		return
	}
	status = resp.StatusCode
	content = string(body)
	return
}

func m_goquery(url string) (body *goquery.Document) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Printf("can not open url: %v\nerr: %v", url, err)
	}
	return doc
}

func main() {
	url := "http://music.163.com/#/discover/toplist"
	//_, status := getUrl(url)
	//fmt.Print(status)
	contents := m_goquery(url)
	fmt.Println(contents)
}
