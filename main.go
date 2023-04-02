package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	//u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	//fmt.Println(u.Scheme)
	//fmt.Println(u.Host)
	//fmt.Println(u.Path)
	//fmt.Println(u.RawQuery)
	//fmt.Println(u.Fragment)
	//
	//fmt.Println(u.IsAbs()) // 絶対URLならtrue
	//fmt.Println(u.Query()) // クエリをマップで取得する
	//
	//url := &url.URL{}
	//url.Scheme = "https:"
	//url.Host = "google.com"
	//q := url.Query()
	//q.Set("q", "Golang") // クエリをセット
	//url.RawQuery = q.Encode()
	//
	//fmt.Println(url)

	//GETメソッド
	res, _ := http.Get("https://example.com")

	fmt.Println(res.StatusCode)

	fmt.Println(res.Proto)

	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])

	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	defer res.Body.Close() // リソースを解放
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body))

	//------------------------------------
	//POSTメソッド

	vs := url.Values{}

	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode()) // => "id=1&message=%E3%83%A1%E3%83%83%E3%82%BB%E3%83@<dtp>{lb}%BC%E3%82%B8"

	res, err := http.PostForm("https://example.com/", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Print(string(body))

}
