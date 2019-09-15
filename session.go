/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 下午12:08
* @Description: 自研 多兼容 session  (设计思路一个中央控制器   一个interface实现控制器下的方法)
* @Github: https://github.com/dollarkillerx
* */
package easyutils

import (
	"errors"
	"github.com/dollarkillerx/easyutils/gcache"
	"time"
)

// session node
type Session struct {
	Name           string
	Identification string                 // 唯一身份标识
	Data           map[string]interface{} // 存储器
	CreationTime   int64 // 创建时间
	ExpirationTime int64 // 过期时间
}

type SessionInterface interface {
	Get(id string) (*Session, error)   // 拥有通过id获取session
	Set(data *Session) (string, error) // 返回id and 错误信息
	SetTime(data *Session,tim time.Duration) (string, error) // 返回id and 错误信息
	Del(id string) error
	Expired(id string) bool   // 检测是否过期  过期返回false 反之true
}


// ================================不同数据源的实现=====================================
// 系统自带 gocache 存储
type GoSessionNode struct {

}


func SessionGetByGoCache() SessionInterface {
	node := GoSessionNode{}
	return &node
}


func (g *GoSessionNode) Get(id string) (*Session,error) {
	// 存储器
	get, b := gcache.CacheGet(id)
	if !b {
		return nil,errors.New("data not ex")
	}

	s,ok := get.(*Session)
	if ok {
		return s,nil
	}

	return nil,errors.New("data not ex")
}

func (g *GoSessionNode) Set(data *Session) (string,error) {
	// 生成随机key
	key := SuperRand()
	err := gcache.CacheSetTime(key, data, 60*60*6)
	if err != nil {
		return "",err
	}
	return key, nil
}

func (g *GoSessionNode) SetTime(data *Session,tim time.Duration) (string,error) {
	// 生成随机key
	key := SuperRand()
	err := gcache.CacheSetTime(key, data, 60*60*tim)
	if err != nil {
		return "",err
	}
	return key, nil
}

func (g *GoSessionNode) Expired(id string) bool {
	exit := gcache.Exit(id)
	if exit {
		return true
	}
	return false
}


func (g *GoSessionNode) Del(id string) error {
	err := gcache.CacheDle(id)
	if err != nil  {
		return err
	}
	return nil
}