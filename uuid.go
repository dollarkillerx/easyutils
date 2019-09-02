package easyutils

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func NewUUID() (string, error) {
	out, err := exec.Command("uuidgen").Output()

	oot := fmt.Sprintf("%s", out)
	oot = strings.TrimSpace(oot)
	return oot, err
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
	body := rand.Intn(999999)
	footer := int(time.Now().UnixNano())

	encode := Sha256Encode(strconv.Itoa(head + body + footer))

	return encode
}