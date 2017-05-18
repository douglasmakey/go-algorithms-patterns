package main

import (
	"flag"
	"net/http"
	"net/url"
	"time"
)

var (
	nurl *url.URL
	u    = flag.String("url", "http://127.0.0.1:3001/", "Connection URL")
	//If you will increase ConnectionCount, validate the maximum number of open file descriptors and change it.
	c    = flag.Int("c", 5000, "Connection Count")
)

func RequestConnect() {
	req := http.Request{
		Method: "GET",
		Header: http.Header{},
		URL:    nurl,
	}
	for {
		conn := http.Client{}
		resp, err := conn.Do(&req)
		if err != nil {
			panic(err.Error())
		}
		resp.Close = true
		resp.Body.Close()
	}
}

func main() {
	flag.Parse()
	nurl, _ = url.Parse(*u)
	for i := 0; i < *c; i++ {
		go RequestConnect()
	}
	time.Sleep(1e10)
}