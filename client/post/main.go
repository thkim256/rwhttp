package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"v1", "v2"},
		"key":  {"vv"},
	}
	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}
