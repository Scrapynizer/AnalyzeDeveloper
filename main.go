package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

var page = "http://www.allitebooks.com/"

func main() {
	content := getContent(page)

	document, e := goquery.NewDocumentFromReader(strings.NewReader(content))

	if e != nil {
		panic(e)
	}

	books := document.Find("article")

	books.Each(func(i int, s *goquery.Selection) {
		header := s.Find("header")
		name := header.Find("h5 a").Text()

		fmt.Println(name)
	})
}

func getContent(url string) string {
	cacheDirectory := "var/cache"
	fileName := cacheDirectory + "/" + getMD5Hash(url)
	fileContent, e := ioutil.ReadFile(fileName)

	if e == nil {
		return string(fileContent)
	}

	response, e := http.Get(url)
	if e != nil {
		panic(e)
	}
	defer response.Body.Close()

	content, e := ioutil.ReadAll(response.Body)

	if e != nil {
		panic(e)
	}

	contentString := string(content)

	ioutil.WriteFile(fileName, content, 0777)

	return contentString
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
