/**
 * @Author: DollarKiller
 * @Description: 压缩字符串
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:38 2019-09-25
 */
package compression

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

type Str struct {
}

func NewStrZip() *Str {
	return &Str{}
}

func (s *Str) Zip(str string) string {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(str)); err != nil {
		return ""
	}
	if err := gz.Flush(); err != nil {
		return ""
	}
	if err := gz.Close(); err != nil {
		return ""
	}
	strc := base64.StdEncoding.EncodeToString(b.Bytes())
	return strc
}

func (s *Str) Unzip(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	rdata := bytes.NewReader(data)
	rc, err := gzip.NewReader(rdata)
	if err != nil {
		return ""
	}
	all, err := ioutil.ReadAll(rc)
	if err != nil {
		return ""
	}
	return string(all)
}
