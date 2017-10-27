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
	_ "reflect"
	"encoding/json"
	"crypto/aes"
	"os"
	"crypto/cipher"
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

//获取歌曲热度评论
func get_hot_comment(id string){
	comment_url := "http://music.163.com/weapi/v1/resource/comments/R_SO_4_" + id + "/?csrf_token="
	params := get_params()

	fmt.Println(comment_url, params)
	return
}

var forth_param string = "0CoJUm6Qyw8W8jud"
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func get_params() (h_encText string){
	iv := commonIV
	first_key := forth_param
	second_key := 16 * 'F'

	//创建加密算法
	c, err := aes.NewCipher([]byte(first_key))
    if err != nil {
        fmt.Printf("Error: NewCipher(%d bytes) = %s", len(first_key), err)
        os.Exit(-1)
    }
	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, iv)
    ciphertext := make([]byte, len(plaintext))
    cfb.XORKeyStream(ciphertext, plaintext)
    fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	h_encText := aes.NewCipher(first_param, first_key, iv)
	h_encText := AES_encrypt(h_encText, second_key, iv)
	return
}

//获取歌单id
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

//获取歌单中歌曲id
func getSongId() (slice []string){
	m_url := "http://music.163.com/api/playlist/detail?id=965267769&updateTime=-1"
	doc, err := goquery.NewDocument(m_url)
	if err != nil{
		fmt.Println(err)
	}
	Html := doc.Text()
	//fmt.Println(Html)

	SongIdJsonConvertMap(Html)
	return
}

//歌单id json数据转换map
func SongIdJsonConvertMap(jsonString string) (SongName, SongId string){
	var SongMapInfo map[string]interface{}

	SongJson := []byte(jsonString)
	if err := json.Unmarshal(SongJson, &SongMapInfo); err != nil{
		panic(err)
	}
	//fmt.Println(SongMapInfo)

	//get slice for song info
	lenTracks := len(SongMapInfo["result"].(map[string]interface{})["tracks"].([]interface{}))

	for i := 0; i < lenTracks; i++{
		SongName := SongMapInfo["result"].(map[string]interface{})["tracks"].([]interface{})[i].(map[string]interface{})["name"]
		floatId := SongMapInfo["result"].(map[string]interface{})["tracks"].([]interface{})[i].(map[string]interface{})["id"].(float64)
		SongId := int(floatId)

		fmt.Println(SongName, SongId)
	}

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
