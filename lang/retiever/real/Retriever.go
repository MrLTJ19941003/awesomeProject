package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get (url string) string{
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println("error : = " , err)
		panic(err)
	}
	result,err := httputil.DumpResponse(resp,true)
	if err != nil {
		panic(err)
	}
	return string(result)
}




