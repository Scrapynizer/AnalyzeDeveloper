package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(getContent("https://github.com"))
}

func getContent(url string) string {
	response, e := http.Get(url)
	defer response.Body.Close()

	if e != nil {
		panic(e)
	}

	content, e := ioutil.ReadAll(response.Body)

	if e != nil {
		panic(e)
	}

	return string(content)
}
