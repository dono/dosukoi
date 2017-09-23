package main

import (
	"strings"
	"testing"
)

func TestSuccessHparser(t *testing.T) {
	// Normal request
	str1 := strings.Join([]string{
		"GET /hoge/fuga HTTP/1.1",
		"Host: newgame.work",
		"Authorization: Basic aG9nZTpmdWdh",
		"Proxy-Authorization: Basic aG9nZTpmdWdh",
		"User-Agent: curl/7.54.0",
		"Accept: image/webp,image/apng,image/*,*/*;q=0.8",
		"Referer: http://google.com",
		"Accept-Encoding: gzip, deflate",
		"Accept-Language: ja-JP,ja;q=0.8,en-US;q=0.6,en;q=0.4",
		"Proxy-Connection: Keep-Alive",
		"",
	}, "\r\n")

	hreq, err := hparser(str1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Success data 1")
	t.Logf("%+v\n", hreq)

	// Don't have Blank-Line
	str2 := strings.Join([]string{
		"GET / HTTP/1.1",
		"Host: example.com",
	}, "\r\n")

	hreq, err = hparser(str2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Success data 2")
	t.Logf("%+v\n", hreq)

	// Be interrupted, should be ignored
	str3 := strings.Join([]string{
		"GET http://example.com/hoge/fuga HTTP/1.0",
		"Host: example.com",
		"Authoriz",
	}, "\r\n")

	hreq, err = hparser(str3)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Success data 3")
	t.Logf("%+v\n", hreq)
}

func TestFailHparser(t *testing.T) {
	// empty
	str1 := ""
	_, err := hparser(str1)
	if err != nil {
		t.Logf("Fail error 1")
		t.Log(err)
		return
	}

	// not HTTP-Request
	str2 := "hogefugahogefuga"
	_, err = hparser(str2)
	if err != nil {
		t.Logf("Fail error 2")
		t.Log(err)
		return
	}

	t.Fatal("Fail-Data was judged to Success-Data")
}
