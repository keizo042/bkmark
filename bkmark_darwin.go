package bkmark

import (
	"os"
)

var (
	ChromeBookMarkFile string = os.Getenv("HOME") + "/Library/Application Support/Google/Chrome/Default/Bookmarks"
)
