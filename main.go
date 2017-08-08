package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, e := http.Get("https://github.com")
	defer response.Body.Close()

	if e != nil {
		panic(e)
	}

	content, e := ioutil.ReadAll(response.Body)

	if e != nil {
		panic(e)
	}

	fmt.Println(string(content))
}
