package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/md5"
	"encoding/hex"
)

func main() {
	fmt.Println(getContent("https://github.com"))
}

func getContent(url string) string {
	cacheDirectory := "var/cache";
	fileName := cacheDirectory + "/" + getMD5Hash(url)
	fileContent, e := ioutil.ReadFile(fileName);

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
