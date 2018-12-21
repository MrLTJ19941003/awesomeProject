package fetcher

import (
	"fmt"
	"testing"
)

func TestLoginRequest(t *testing.T) {
	url := "https://passport.jiayuan.com/dologin.php?pre_url=http://www.jiayuan.com/usercp/"
	fromparams := make(map[string]interface{})
	fromparams["name"] = "19928709269"
	fromparams["password"] = "liu01234"
	fromparams["remem_pass"] = ""
	//reuqest := models.CreateReuqest(method, url, fromparams)
	cookie, _ := LoginRequest(fromparams, url)
	for _,v := range cookie{
		fmt.Printf("%v \n",v)
	}

}
