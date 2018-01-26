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

func mGoquery(mUrl string) {
	doc, err := goquery.NewDocument(mUrl)
	if err != nil {
		fmt.Printf("can not open url: %v\nerr: %v", mUrl, err)
	}
	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
		if s.Find(".thumb").Nodes == nil && s.Find(".video_holder").Nodes == nil {
			content := s.Find(".content").Text()
			fmt.Printf("%s", content)
		}
	})
	//return doc
}

func music163Search(s, mUrl string) (content string) {
	resp, err := http.PostForm(mUrl, url.Values{"s": {s}, "type": {"100"}, "offset": {"0"}, "limit": {"1"}})

	resp.Header.Set("Referer", "http://music.163.com/")
	//fmt.Printf("Header info: %v\n\n", resp.Header)

	if err != nil {
		fmt.Printf("url: %v\n post err: %v", mUrl, err)
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
	mUrl := "http://music.163.com/discover/playlist/?order=hot&cat=全部&limit=5&offset=1"

	titleMap := make(map[string]string)
	//id_slice := make([]string, 0)
	re := regexp.MustCompile("\\?id=")

	doc, err := goquery.NewDocument(mUrl)
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
		songId := re.Split(link, 2)
		//id_slice = append(id_slice, song_id[1])

		title, ok := s.Find("a").Attr("title")
		if !ok {
			fmt.Printf("SongList title get err: %v", ok)
		}
		titleMap[title] = songId[1]
	})
	return titleMap
}

//获取歌单中歌曲id
func getSongId(songid string) (map[string]float64) {
	//fmt.Println(songid)
	mUrl := "http://music.163.com/api/playlist/detail?id=" + songid + "&updateTime=-1"
	doc, err := goquery.NewDocument(mUrl)
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

func FloatToString(inputNum float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(inputNum, 'f', -1, 64)
}

//获取歌曲热度评论
func getHotComment(id float64) {
	commentUrl := "http://music.163.com/weapi/v1/resource/comments/R_SO_4_" + FloatToString(id) + "?csrf_token="
	fmt.Println(commentUrl)
	params, err := getParams()
	fmt.Println(params)
	encSecKey := "257348aecb5e556c066de214e531faadd1c55d814f9be95fd06d6bff9f4c7a41f831f6394d5a3fd2e3881736d94a02ca919d952872e7d0a50ebfa1769a7a62d512f5f1ca21aec60bc3819a9c3ffca5eca9a0dba6d6f7249b06f5965ecfff3695b54e1c28f3f624750ed39e7de08fc8493242e26dbc4484a01c76f739e135637c"

	resp, err := http.PostForm(commentUrl, url.Values{"params": {params}, "encSecKey": {encSecKey}})
	resp.Header.Set("Cookie", "appver=1.5.0.75771;")
	resp.Header.Set("Referer", "http://music.163.com/")

	if err != nil{
		fmt.Printf("get_hot_comment: url: %v\npost err: %v", commentUrl, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get_hot_comment can not get url body.err: %v\n", err)
	}

	content := string(body)
	fmt.Println(content)
	//commentJsonProcess(content)
}

//python获取评论url加密参数
func getParams() (string, error){
	PythonBin := "/usr/bin/python"
	cmd := exec.Command(PythonBin, "-c", "import src.encrypt_params as en; print(en.get_params())")
	stdout, err := cmd.CombinedOutput()
	if err != nil{
		return "", fmt.Errorf("exec python err: %s", err)
	}
	return string(stdout), nil
}

//评论json数据转换
func commentJsonProcess(comment string){
	var commMap map[string]interface{}

	commJson := []byte(comment)
	if err := json.Unmarshal(commJson, &commMap); err != nil{
		fmt.Errorf("comment err: %s", err)
	}

	commentTotal := commMap["total"].(float64)
	if commentTotal > 100000 {
		//fmt.Println(comment_total)
		hotCounts := len(commMap["hotComments"].([]interface{}))

		for i := 0; i < hotCounts; i++ {
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
			getHotComment(id)
		}
	}
}

func main() {
	//MusicCommentSpider()

	//a := getSongListId()
	//fmt.Println(a)

	// get song id
	//b := getSongId("751345753")
	//fmt.Println(b)

	//加密参数
	//getParams()

	getHotComment(109998)
}
