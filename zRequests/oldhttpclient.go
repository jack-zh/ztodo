// +build nobuild

package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HTTPClientGet(url string) (string, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func HTTPClientPost(url string, task string) (string, error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("task="+task))

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func HTTPClientPostForm(url string) (string, error) {
	resp, err := http.PostForm(url,
		url.Values{"key": {"Value"}, "id": {"123"}})

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func HTTPClientDo(url string, dotype string) (string, error) {
	client := &http.Client{}

	// dotype "POST" GET DELETE PUT ...
	req, err := http.NewRequest(dotype, url, strings.NewReader("name=cjb"))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	req.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "chrome 100")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
