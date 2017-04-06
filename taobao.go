package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/anlgo/mahonia"
)

// Taobao is ...
func Taobao() {
	_url := "https://item.taobao.com/item.htm?spm=a219r.lm872.14.60.4luJFW&id=528330539743&ns=1&abbucket=9"

	doc, _ := goquery.NewDocument(_url)

	_title := doc.Find("#J_sidebar_config").Next().Text()

	fmt.Println(_title)

	_js := GBKToUtf(_title)

	fmt.Println("GBK to UTF-8: ", _js)

	JsRun(_js)

}

// GBKToUtf is ...
func GBKToUtf(str string) string {

	dec := mahonia.NewDecoder("gbk")
	ret, ok := dec.ConvertStringOK(str)

	if !ok {
		ret = ""
	}

	return ret

}
