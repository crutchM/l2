package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var a string
	fmt.Scan(&a)
	resp, err := http.Get(a)
	defer resp.Body.Close()
	if err != nil {
		panic("invalid url")
	}
	record("./site.html", resp)
}

func record(file string, resp *http.Response) {
	res, err := os.Create(file)
	defer res.Close()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	res.WriteString(string(body))
}
