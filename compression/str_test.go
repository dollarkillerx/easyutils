/**
 * @Author: DollarKillerX
 * @Description: str_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:18 2019/11/22
 */
package compression

import (
	"log"
	"testing"
)

func TestStr_Zip(t *testing.T) {
	tagText := ""

	zip := NewStrZip()
	s := zip.Zip(tagText)
	log.Println(s)
	unzip := zip.Unzip(s)
	if unzip != tagText {
		t.Fatal("eeee")
	}

	a1 := len(tagText)
	a2 := len(s)
	c := a1 - a2
	if a1 > a2 {
		log.Println("通过")
		log.Printf("压缩率: %d", c)
	} else {
		log.Println("不通过")
		log.Printf("压缩率: %d", c)
		log.Println(a1)
		log.Println(a2)
	}
}
