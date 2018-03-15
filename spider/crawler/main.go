package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERR: %s", resp.StatusCode)
		return
	}

	//trasform GBK  TO UTF8
	//utf8Reader :=transform.NewReader(resp.Body,
	//	simplifiedchinese.GBK.NewDecoder())

	e := determinEncoding(resp.Body)
	utf8Reader :=transform.NewReader(resp.Body,
		e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)
	printCityList(all)

}

func printCityList(contents []byte){
	re :=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	//matches := re.FindAll(contents,-1)
	matches := re.FindAllSubmatch(contents,-1)
	for _,m := range matches {
		fmt.Printf("City: %s, URL: %s\n",m[2],m[1])
	}
	fmt.Printf("Matches found : %d\n",len(matches))
}


func determinEncoding(r io.Reader) encoding.Encoding {
	bytes,err :=bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e,_,_ :=charset.DetermineEncoding(bytes,"")

	return e
}