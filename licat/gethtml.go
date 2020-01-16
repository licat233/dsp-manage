package licat

import (
	"fmt"
	"gopacong/common"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

//GetHTML 爬取所需要的数据（消耗额）
func GetHTML(cookie string) {
	url := "https://licat.work/userlist"
	client := &http.Client{}
	var resp *http.Response
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	// client.Jar = jar

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0")
	req.Header.Add("Host", "licat.work")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("TE", "Trailers")
	req.Header.Add("Content-type", "text/html; charset=utf-8")
	req.Header.Add("Cookie", cookie)

	//发送请求数据之服务器，并获取响应数据
	resp, err = client.Do(req)
	time.Sleep(5 * time.Second)

	if err != nil {
		log.Println("Client Do Failed:", err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		os.Exit(1)
	}

	// 打印所有请求头信息
	for k, v := range req.Header {
		fmt.Println("请求头信息 :", k, "=", v)
	}

	// 打印所有响应头信息
	for k, v := range resp.Header {
		fmt.Println("响应头信息 :", k, "=", v)
	}

	// 打印正文信息
	zip, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("ioutil.ReadAll failed:", err)
		os.Exit(1)
	}

	gzipData, _ := common.GzipDecode(zip)
	content := string(gzipData)
	fmt.Println("正文信息:", content)

	//正则公式
	reg := regexp.MustCompile(`\d*%`)
	prob := reg.FindAllString(content,-1)
	fmt.Println(prob)
}
