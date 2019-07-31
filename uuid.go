package easyutils

import (
	"fmt"
	"os/exec"
	"strings"
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
