package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	input := []byte("### topgoer.com是个不错的go文档网站")
	unsafe := blackfriday.MarkdownCommon(input)
	html := bluemonday.UGCPolicy()(unsafe)
	fmt.Println(string(html))
}
