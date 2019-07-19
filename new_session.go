/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-17
* Time: 下午5:58
* */
package easyutils

// 重构
// 采用 micro kernel 设计模式

type Session interface {
	//@生成session
	//@Params time 过期时间 h
	//@return sessionId
	SessionGenerate(time int64) (string, error)

	//@用户设置一些东西
	SessionSet(session_id string, key string, value interface{}) error

	//@验证用户
	SessionCheck(session_id string) bool

	//@删除用户
	SessionDel(session_id string) bool
}
