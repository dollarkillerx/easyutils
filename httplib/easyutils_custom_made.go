package httplib

import (
	"bufio"
	"github.com/dollarkillerx/easyutils"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// httpLib easyUtils定制版

// 假装用户去请求
func EuUserGet(url string) ([]byte, error) {
	get := Get(url).SetUserAgent(easyutils.ReptileGetUserAgent())
	return get.Bytes()
}

// 假装搜索引擎去请求
func EuSpiderGet(url string) ([]byte, error) {
	get := Get(url).SetUserAgent(easyutils.ReptileGetSpiderAgent())
	return get.Bytes()
}

func EuUserGetEncoding(url string) ([]byte, error) {
	bytes, e := EuUserGet(url)
	if e != nil {
		return nil, e
	}
	data := strings.NewReader(string(bytes))
	utf8Reader := transform.NewReader(data, determineencoding(data).NewDecoder())

	//将其他编码的reader转换为常用的utf8reader
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func determineencoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}


// 配置代理下载
func ProxyDow(tagurl,proxy string) (*http.Response,error) {
	response, e := Get(tagurl).SetUserAgent(easyutils.ReptileGetUserAgent()).SetProxy(func(request *http.Request) (tagurl *url.URL, e error) {
		u := new(url.URL)
		u.Scheme = "http"
		u.Host = proxy
		return u, nil
	}).Response()

	return response,e
}