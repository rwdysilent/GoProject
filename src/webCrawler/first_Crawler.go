package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
	_ "go/doc"
	"regexp"
	"reflect"
)

func getUrl(url string) (content string, status int) {
	resp, err := http.Get(url)
	resp.Header.Set("Referer", "defifition referrer info")
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
	fmt.Println(resp.Header)
	return
}

func m_goquery(m_url string){
	doc, err := goquery.NewDocument(m_url)
	if err != nil {
		fmt.Printf("can not open url: %v\nerr: %v", m_url, err)
	}
	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
		if s.Find(".thumb").Nodes == nil && s.Find(".video_holder").Nodes == nil {
            content := s.Find(".content").Text()
            fmt.Printf("%s", content)
        }
    })
	//return doc
}


func music_163_search(s, m_url string) (content string){
	resp, err := http.PostForm(m_url, url.Values{"s": {s}, "type": {"100"}, "offset": {"0"}, "limit": {"1"}})

	resp.Header.Set("Referer", "http://music.163.com/")
	//fmt.Printf("Header info: %v\n\n", resp.Header)

	if err != nil {
		fmt.Printf("url: %v\n post err: %v", m_url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("can not get url body.err: %v\n", err)
	}
	content = string(body)
	fmt.Println(content)
	return
}

func getSongListId() (slice []string){
	m_url := "http://music.163.com/discover/playlist/?order=hot&cat=全部&limit=5&offset=1"

	id_slice := make([]string, 0)
	re := regexp.MustCompile("\\?id=")

	doc, err := goquery.NewDocument(m_url)
	if err != nil {
		log.Fatal(err)
	}

	// Find the song id items
	doc.Find(".m-cvrlst li .dec").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Find("a").Attr("href")
		text := s.Find("a").Text()
		if ok {
			fmt.Printf("info %d: %s - %s\n", i, link, text)
		}
		song_id := re.Split(link, 2)
		id_slice = append(id_slice, song_id[1])
	})

	return id_slice
}

func getSongId() (slice []string){
	m_url := "http://music.163.com/api/playlist/detail?id=965267769&updateTime=-1"
	doc, err := goquery.NewDocument(m_url)
	if err != nil{
		fmt.Println(err)
	}
	Html := doc.Text()
	fmt.Println(reflect.TypeOf(Html))

	return
}

func main() {
	// search api
	//s := ""
	//m_url := "http://music.163.com/api/search/pc"
	//music_163_search(s, m_url)

	// get song list Id
	//a := getSongListId()
	//fmt.Println(a)

	// get song id
	getSongId()
}
