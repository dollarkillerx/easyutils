package easyutils

import (
	"fmt"
	"github.com/dollarkillerx/easyutils/uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func NewUUID() (string, error) {
	v4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	str := fmt.Sprintf("%s", v4)
	return str, nil
}

// 获取没有 - 的uuid
func NewUUIDSimplicity() (string, error) {
	s, e := NewUUID()
	var u string
	for _, k := range s {
		if k != '-' {
			u = fmt.Sprintf("%s%s", u, string(k))
		}
	}
	u = strings.TrimSpace(u)
	return u, e
}

func SuperRand() string {
	head := int(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	body := rand.Intn(999999)
	footer := int(time.Now().UnixNano())

	encode := Sha256Encode(strconv.Itoa(head + body + footer))

	return encode
}
