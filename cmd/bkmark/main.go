package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/keizo042/bkmark"
)

func main() {
	var err error
	//searchByUrl := flag.Bool("url", false, "search by url")
	browserType := flag.String("browser", "chrome", "specify browser")
	var params *bkmark.Params
	params, err = bkmark.LoadBookMark()
	if err != nil {
		fmt.Printf("fail to load browser bookmark: %v\n", err)
		return
	}
	params, err = bkmark.FilterByPeco(params)
	if err != nil {
		fmt.Printf("fail to filter: %v\n", err)
		return
	}
	for _, v := range params.Bookmarks {
		if v.Url == "" {
			continue
		}
		if _, err := url.Parse(v.Url); err != nil {
			fmt.Println(err)
			continue
		}
		if err := bkmark.OpenURL(v.Url); err != nil {
			fmt.Println(err)
		}
	}
}
