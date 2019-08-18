package httplib

import "github.com/dollarkillerx/easyutils"

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
