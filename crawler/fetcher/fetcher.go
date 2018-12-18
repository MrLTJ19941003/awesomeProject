package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const userAgent  = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.75 Safari/537.36"
var rateLimiter = time.Tick(10 * time.Millisecond)
func Fetch(url string)([]byte,error){
	//resp,err := http.Get(url)
	<-rateLimiter
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err!=nil	{
		fmt.Println("获取地址错误")
	}
	req.Header.Add("User-Agent",userAgent)
	resp, err := client.Do(req)

	if err !=nil{
		return nil ,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error : status code ",resp.StatusCode)
		return nil,
			fmt.Errorf("Wrong status code : %d" , resp.StatusCode)
	}
	//urf8Reader := transform.NewReader(resp.Body,
	//simplifiedchinese.GBK.NewDecoder())
	all, err := ioutil.ReadAll(resp.Body)
	return all,err
}
