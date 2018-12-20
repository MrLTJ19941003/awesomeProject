package fetcher

import (
	"awesomeProject/crawler/shijiejiayuan/type"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)


var rateLimiter = time.Tick(10 * time.Millisecond)
//发送一个get请求
func SendRequests(r types.Request, client *http.Client) (string, error) {
	<- rateLimiter
	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		return "", err
	}
	setRequestHeader(req) //设置请求头参数
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "",err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error : %v" ,err)
	}
	contents := string(buf)
	return contents, nil
}

//发送登陆一个请求
func LoginRequest(params map[string]interface{}, client *http.Client,url string) (string, error) {
	var reader *strings.Reader
	var req *http.Request
	var err error
	if params != nil {
		reader = strings.NewReader(fmt.Sprintf("name=%s&password=%s&remem_pass=%s", params["name"].(string), params["password"].(string), params["remem_pass"].(string)))
		req, err = http.NewRequest("POST", url, reader)
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}
	if err != nil {
		return "", err
	}
	setRequestHeader(req) //设置请求头参数
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "",err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error : %v" ,err)
	}
	contents := string(buf)
	return contents, nil
}

//设置请求头
func setRequestHeader(req *http.Request) {
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "utf-8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Origin", "http://center.qianlima.com")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36")
	req.Header.Set("Host", "center.qianlima.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}
