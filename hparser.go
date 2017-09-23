package main

import (
	"bufio"
	"encoding/base64"
	"errors"
	"strings"
)

var (
	methods  = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}
	versions = []string{"HTTP/1.0", "HTTP/1.1"}
)

// check if the target is included in the list
func containArr(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

// parse info from HTTP-Request, and return httpRequest(struct)
func hparser(str string) (httpRequest, error) {

	var hreq httpRequest

	scanner := bufio.NewScanner(strings.NewReader(str))

	// read first-line
	scanner.Scan()
	firstLine := scanner.Text()
	arr := strings.Split(firstLine, " ")
	if len(arr) != 3 {
		return hreq, errors.New("Invalid HTTP-Request-Line, wrong format")
	}
	hreq.Method = arr[0]
	path := arr[1]
	hreq.Version = arr[2]
	if !containArr(methods, hreq.Method) || !containArr(versions, hreq.Version) {
		return hreq, errors.New("Invalid HTTP-Request-Line, wrong value")
	}

	// read second-line ~ final-line
	for scanner.Scan() {
		line := scanner.Text()
		hdr := strings.SplitN(line, ":", 2)
		if len(hdr) != 2 {
			continue
		}

		title := hdr[0]
		content := strings.TrimSpace(hdr[1])

		switch title {
		case "Host":
			host := content
			if strings.HasPrefix(path, "http://") {
				hreq.URL = path
				hreq.UseProxy = true
			}
			hreq.URL = "http://" + host + path
		case "Authorization":
			b64arr := strings.SplitN(content, " ", 2)
			if len(b64arr) == 2 && b64arr[0] == "Basic" {
				data, err := base64.StdEncoding.DecodeString(b64arr[1])
				if err != nil {
					continue
				}
				hreq.BasicAuth = string(data)
			}
		case "Proxy-Authorization":
			b64arr := strings.SplitN(content, " ", 2)
			if len(b64arr) == 2 && b64arr[0] == "Basic" {
				data, err := base64.StdEncoding.DecodeString(b64arr[1])
				if err != nil {
					continue
				}
				hreq.ProxyAuth = string(data)
			}
		case "Referer":
			hreq.Referer = content
		case "User-Agent":
			hreq.UserAgent = content
		}
	}

	return hreq, nil
}
