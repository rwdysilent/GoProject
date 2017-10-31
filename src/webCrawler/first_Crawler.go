package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
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
	"io"
	_"strings"
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

//获取歌曲热度评论
func get_hot_comment(id string) (string) {
	comment_url := "http://music.163.com/weapi/v1/resource/comments/R_SO_4_" + id + "/?csrf_token="
	fmt.Println(comment_url)
	params := "O5/yxckUkfK03FP34r7bgJVnmX5k2/G/l+JCIrgOQwzIIaSLS4Whg5hM1NqjOg7Q8fUC73m3WAcoCTPlNUrAVcycdW/bHEENz/Od0HbqY48y98a5kdtGtQCEDPQo2J5G"
	encSecKey := "257348aecb5e556c066de214e531faadd1c55d814f9be95fd06d6bff9f4c7a41f831f6394d5a3fd2e3881736d94a02ca919d952872e7d0a50ebfa1769a7a62d512f5f1ca21aec60bc3819a9c3ffca5eca9a0dba6d6f7249b06f5965ecfff3695b54e1c28f3f624750ed39e7de08fc8493242e26dbc4484a01c76f739e135637c"

	resp, err := http.PostForm(comment_url, url.Values{"params": {params}, "encSecKey": {encSecKey}})
	resp.Header.Set("Cookie", "appver=1.5.0.75771;")
	resp.Header.Set("Referer", "http://music.163.com/")

	if err != nil{
		fmt.Printf("url: %v post err: %v", comment_url, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("can not get url body.err: %v\n", err)
	}

	fmt.Println()
	content := string(body)
	fmt.Println(content)
	return content
}

func get_params() (string) {
	plaintext := []byte("{rid:\"\", offset:\"0\", total:\"true\", limit:\"2\", csrf_token:\"\"}")
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	commonIV := []byte("0102030405060708")
	key := []byte("0CoJUm6Qyw8W8jud")

	//创建加密算法
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
		//os.Exit(-1)
	}

	//加密字符串
	//ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	ciphertext := make([]byte, len(plaintext))
	//commonIV := ciphertext[:aes.BlockSize]

	cfb := cipher.NewCBCEncrypter(c, commonIV)
	//cfb.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	cfb.CryptBlocks(ciphertext, plaintext)

	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	return string(ciphertext)
}

func get_params_example(){
	key := []byte("0CoJUm6Qyw8W8jud")
	plaintext := []byte("exampleplaintext")
	//plaintext := []byte("{rid:\"\", offset:\"0\", total:\"true\", limit:\"2\", csrf_token:\"\"}")

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.

	//pad := 16 - len(plaintext) % 16
    //plaintext = []byte(string(plaintext) + strings.Repeat(string(pad), pad))

	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	fmt.Printf("%x\n", ciphertext)
}

//获取歌单id
func getSongListId() (slice []string) {
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
func getSongId() (slice []string) {
	m_url := "http://music.163.com/api/playlist/detail?id=965267769&updateTime=-1"
	doc, err := goquery.NewDocument(m_url)
	if err != nil {
		fmt.Println(err)
	}
	Html := doc.Text()
	//fmt.Println(Html)

	SongIdJsonConvertMap(Html)
	return
}

//歌单id json数据转换map
func SongIdJsonConvertMap(jsonString string) (SongName, SongId string) {
	var SongMapInfo map[string]interface{}

	SongJson := []byte(jsonString)
	if err := json.Unmarshal(SongJson, &SongMapInfo); err != nil {
		panic(err)
	}
	//fmt.Println(SongMapInfo)

	//get slice for song info
	lenTracks := len(SongMapInfo["result"].(map[string]interface{})["tracks"].([]interface{}))

	for i := 0; i < lenTracks; i++ {
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
	//getSongId()

	//加密参数
	//get_params()
	get_params_example()

	//get_hot_comment("30953009")
}
