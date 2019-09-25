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

func (s *Str) Zip(str string) string {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(str)); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	strc := base64.StdEncoding.EncodeToString(b.Bytes())
	return strc
}

func (s *Str) Unzip(str string) string {
	data, _ := base64.StdEncoding.DecodeString(str)
	rdata := bytes.NewReader(data)
	rc, _ := gzip.NewReader(rdata)
	all, _ := ioutil.ReadAll(rc)
	return string(all)
}