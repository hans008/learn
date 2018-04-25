package real

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type Retriver struct {
	UserAgent string
	Timeout time.Duration
}

func (r Retriver) Get(url string) string {
	resp,err := http.Get(url)
	if err != nil{
		panic(resp)
	}
	result,err := httputil.DumpResponse(resp,true)
	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}

