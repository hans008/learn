package main

import (
	"hans/learn/other/retriever/real"
	"fmt"
	"hans/learn/other/retriever/mock"
)



//定义了一个接口Retriver
type Retriver interface {
	Get(url string) string
}


//使用者download
func download(r Retriver) string {
	return r.Get("http://www.baidu.com")
}



func main(){
	var r Retriver
	r = real.Retriver{}
	inspect(r)
	//断言测试
	if mockRetriever ,ok := r.(mock.Retriever);ok{
		fmt.Println(mockRetriever.Contents)
	}else{
		fmt.Println("not a mock retriever")
	}
}


func inspect (r Retriver){
	fmt.Printf("%T %v\n",r,r)
	switch v :=r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *real.Retriver:
		fmt.Println("Contents:",v.UserAgent)
	}

}

