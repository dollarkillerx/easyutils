package utils

import (
	"strings"
	"testing"
)

func TestNewUUID(t *testing.T) {
	s, e := NewUUID()
	if e!=nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestNewUUIDSimplicity(t *testing.T) {
	s, e := NewUUIDSimplicity()
	if e!=nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestMd5String(t *testing.T) {
	s := Md5Encode("123")
	s = strings.ToUpper(s)
	if s != "202CB962AC59075B964B07152D234B70" {
		t.Error("非正常",s)
	}
}

func TestGetCurrentTime(t *testing.T) {
	time := GetCurrentTime()
	t.Log(time)
}

func TestGetTimeToString(t *testing.T) {
	s, e := GetTimeToString("1558057058")
	if e!=nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestGetTimeStringToTime(t *testing.T) {
	s, e := GetTimeStringToTime("2019-05-17")
	if e!=nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestGetSession(t *testing.T) {
	session := GetSession("dollarkiller")
	t.Log(session)
	bool := CheckSession(session)
	t.Log(SessionMap.Load(session))
	t.Log(bool)
}