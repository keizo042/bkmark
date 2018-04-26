package bkmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Conf struct {
}

type Bookmark struct {
	Name string
	Url  string
}

type ChromeBookmarkToml struct {
	Roots   RootsToml
	Version int
}

type RootsToml struct {
	BookmarkBar BookmarkBarToml `json:"bookmark_bar"`
}

type BookmarkBarToml struct {
	Children []BookmarkToml
}

type BookmarkToml struct {
	Name string
	Id   string
	Type string
	// url
	Url string
	// folder
	Children []BookmarkToml `json:"children"`
}

type Params struct {
	Bookmarks map[string]Bookmark
}

var DefaultConf = Conf{}

func LoadBookMark() (*Params, error) {
	var v ChromeBookmarkToml
	f, err := os.Open(ChromeBookMarkFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&v); err != nil {
		return nil, err
	}
	bookmarks := fixBookmark(v.Roots.BookmarkBar.Children)
	if bookmarks == nil {
		return nil, fmt.Errorf("chrome bookmark is empty")
	}
	return &Params{
		Bookmarks: bookmarks,
	}, nil
}

func fixBookmark(bookmarkToml []BookmarkToml) map[string]Bookmark {
	bookmark := make(map[string]Bookmark)
	for _, b := range bookmarkToml {
		switch b.Type {
		case "url":
			bookmark[b.Name] = Bookmark{
				Name: b.Name,
				Url:  b.Url,
			}
		case "folder":
			b2 := fixBookmark(b.Children)
			for k, v := range b2 {
				bookmark[b.Name+"/"+k] = v
			}
		default:
			continue
		}
	}
	return bookmark
}

func FilterByPeco(p *Params) (*Params, error) {
	var b string
	for k, _ := range p.Bookmarks {
		b += k + "\n"
	}
	cmd := exec.Command("peco")
	cmd.Stdin = bytes.NewBuffer([]byte(b))
	byteFilted, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	bookmark := make(map[string]Bookmark)
	for _, name := range strings.Split(string(byteFilted), "\n") {
		bookmark[name] = p.Bookmarks[name]
	}
	return &Params{
		Bookmarks: bookmark,
	}, nil
}

func OpenURL(url string) error {
	cmd := exec.Command("open", url)
	return cmd.Run()
}
