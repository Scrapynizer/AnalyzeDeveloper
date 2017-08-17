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

type Link struct {
	Name string
	Url  string
}

type Book struct {
	Title   Link
	Authors []Link
}

func main() {
	content := getContent(page)

	document, e := goquery.NewDocumentFromReader(strings.NewReader(content))

	if e != nil {
		panic(e)
	}

	books := []Book{}

	bookListSelection := document.Find("article")

	bookListSelection.Each(func(i int, bookSelection *goquery.Selection) {
		headerSelection := bookSelection.Find("header")
		titleSelection := headerSelection.Find("h2 a")

		authors := []Link{}
		authorListSelection := headerSelection.Find("h5 a")
		authorListSelection.Each(func(i int, authorSelection *goquery.Selection) {
			authorUrl, _ := authorSelection.Attr("href")
			authors = append(authors, Link{
				authorSelection.Text(),
				authorUrl,
			})
		})

		titleUrl, _ := titleSelection.Attr("href")
		titleLink := Link{
			titleSelection.Text(),
			titleUrl,
		}
		books = append(books, Book{
			titleLink,
			authors,
		})
	})

	fmt.Print(books)
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
