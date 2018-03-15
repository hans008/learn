package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"log"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	e := determinEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,
		e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determinEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
