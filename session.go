/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 下午12:08
* */
package utils

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// session 库
var (
	SessionMap sync.Map
)

type SessionNode struct {
	Name string
	CreationTime int64 // 创建时间
	ExpirationTime int64 // 过期时间
}

// 获得session
func GetSession(name string) string {
	timeNano := time.Now().UnixNano()
	time := time.Now().Unix()
	outtime := time + 6*60*60
	intn := rand.Intn(100000)
	encode := Md5Encode(strconv.FormatInt(timeNano, 10) + strconv.Itoa(intn))
	node := &SessionNode{
		Name:         name,
		CreationTime: time,
		ExpirationTime: outtime,
	}

	SessionMap.Store(encode,node)
	return encode
}

// 验证session
func CheckSession(sessionId string) bool {
	if sessionId == "" || len(sessionId) == 0 {
		return false
	}
	value, ok := SessionMap.Load(sessionId)
	if ok != true {
		return false
	}

	node := value.(*SessionNode)
	nowTime := time.Now().Unix()
	if nowTime >= node.CreationTime && nowTime < node.ExpirationTime {
		return true
	}
	return false
}


