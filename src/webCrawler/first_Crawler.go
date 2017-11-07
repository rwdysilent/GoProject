package main

import (
	_ "crypto/aes"
	_ "crypto/cipher"
	_ "crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "go/doc"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	_"os"
	_ "reflect"
	"regexp"
	_ "io"
	_"strings"
	_ "strings"
	"os/exec"
	"strconv"
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

func m_goquery(m_url string) {
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

func music_163_search(s, m_url string) (content string) {
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


//获取歌单id
func getSongListId() (map[string]string) {
	m_url := "http://music.163.com/discover/playlist/?order=hot&cat=全部&limit=5&offset=1"

	title_map := make(map[string]string)
	//id_slice := make([]string, 0)
	re := regexp.MustCompile("\\?id=")

	doc, err := goquery.NewDocument(m_url)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(doc.Html())

	// Find the song id items
	doc.Find(".m-cvrlst li .dec").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Find("a").Attr("href")
		if !ok {
			fmt.Printf("SongListId find err: %v", ok)
		}
		song_id := re.Split(link, 2)
		//id_slice = append(id_slice, song_id[1])

		title, ok := s.Find("a").Attr("title")
		if !ok {
			fmt.Printf("SongList title get err: %v", ok)
		}
		title_map[title] = song_id[1]
	})
	return title_map
}

//获取歌单中歌曲id
func getSongId(songid string) (map[string]float64) {
	//fmt.Println(songid)
	m_url := "http://music.163.com/api/playlist/detail?id=" + songid + "&updateTime=-1"
	doc, err := goquery.NewDocument(m_url)
	if err != nil {
		fmt.Printf("getSongId err: %v", err)
	}
	Html := doc.Text()

	//获取歌曲名与对应id
	SongInfoMap := SongIdJsonConvertMap(Html)
	return SongInfoMap
}

//歌单id json数据, 返回包含歌曲名与id的map
func SongIdJsonConvertMap(jsonString string) (map[string]float64) {
	var SongMapInfo map[string]interface{}
	FilterSongInfo := make(map[string]float64)

	SongJson := []byte(jsonString)
	if err := json.Unmarshal(SongJson, &SongMapInfo); err != nil {
		panic(err)
	}
	//fmt.Println(SongMapInfo)

	//get slice for song info
	lenTracks := len(SongMapInfo["result"].(map[string]interface{})["tracks"].([]interface{}))

	for i := 0; i < lenTracks; i++ {
		SongName := SongMapInfo["result"].(map[string]interface{})["tracks"].([]interface{})[i].(map[string]interface{})["name"].(string)
		floatId := SongMapInfo["result"].(map[string]interface{})["tracks"].([]interface{})[i].(map[string]interface{})["id"].(float64)
		//SongId := int(floatId)

		FilterSongInfo[SongName] = floatId
		//fmt.Println(FilterSongInfo)
	}
	return FilterSongInfo
}

func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}

//获取歌曲热度评论
func get_hot_comment(id float64) {
	comment_url := "http://music.163.com/weapi/v1/resource/comments/R_SO_4_" + FloatToString(id) + "/?csrf_token="
	//fmt.Println(comment_url)
	params := get_params()
	encSecKey := "257348aecb5e556c066de214e531faadd1c55d814f9be95fd06d6bff9f4c7a41f831f6394d5a3fd2e3881736d94a02ca919d952872e7d0a50ebfa1769a7a62d512f5f1ca21aec60bc3819a9c3ffca5eca9a0dba6d6f7249b06f5965ecfff3695b54e1c28f3f624750ed39e7de08fc8493242e26dbc4484a01c76f739e135637c"

	resp, err := http.PostForm(comment_url, url.Values{"params": {params}, "encSecKey": {encSecKey}})
	resp.Header.Set("Cookie", "appver=1.5.0.75771;")
	resp.Header.Set("Referer", "http://music.163.com/")

	if err != nil{
		fmt.Printf("get_hot_comment: url: %v\npost err: %v", comment_url, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get_hot_comment can not get url body.err: %v\n", err)
	}

	content := string(body)
	commentJsonProcess(content)
}

//python获取评论url加密参数
func get_params() (string){
	cmd := exec.Command("C:/Python27/python", "-c", "import src.encrypt_params as en; print(en.get_params())")
	stdout, err := cmd.CombinedOutput()
	if err != nil{
		fmt.Println(err)
	}
	return string(stdout)
}

//评论json数据转换
func commentJsonProcess(comment string){
	var commMap map[string]interface{}

	commJson := []byte(comment)
	if err := json.Unmarshal(commJson, &commMap); err != nil{
		panic(err)
	}

	comment_total := commMap["total"].(float64)
	if comment_total > 100000 {
		//fmt.Println(comment_total)
		hot_counts := len(commMap["hotComments"].([]interface{}))

		for i := 0; i < hot_counts; i++ {
			username := commMap["hotComments"].([]interface{})[i].(map[string]interface{})["user"].(map[string]interface{})["nickname"].(string)
			comment := commMap["hotComments"].([]interface{})[i].(map[string]interface{})["content"].(string)
			fmt.Printf("\n%v: %v\n\n", username, comment)
		}
	}
}

func MusicCommentSpider(){
	SongListId := getSongListId()
	for title, listId := range SongListId {
		fmt.Println(title)
		SongInfo := getSongId(string(listId))
		for name, id := range SongInfo {
			fmt.Printf("Song name : %v\n", name)
			get_hot_comment(id)
		}
	}
}

func main() {
	MusicCommentSpider()

	//a := getSongListId()
	//fmt.Println(a)

	// get song id
	//getSongId()

	//加密参数
	//get_params()

	//get_hot_comment(30953009)
}
