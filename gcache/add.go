/**
 * @Author: DollarKiller
 * @Description: 添加方法
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:26 2019-09-15
 */
package gcache

import "time"

var (
	gc Cache
)

func init() {
	gc = New(20).
		LRU().
		Build()

}

// 返回cache
//func GCache() Cache {
//
//	return gc
//}

// 获取
func CacheGet(key interface{}) (interface{},bool) {
	get, e := gc.Get(key)
	if e != nil {
		return "",false
	}

	if get == "" {
		return "",false
	}

	return get,true
}

// 设置 有过期时间
func CacheSetTime(key,data interface{},tim time.Duration) error {
	err := gc.SetWithExpire(key, data, tim)
	return err
}

// 设置
func CacheSet(key,data interface{}) error {
	err := gc.Set(key, data)
	return err
}

// 删除
func CacheDle(key interface{}) error {
	err := gc.SetWithExpire(key, "", time.Microsecond)
	return err
}

// 检测是否存在
func Exit(key interface{}) bool {
	get, e := gc.Get(key)
	if e != nil {
		return false
	}

	if get == "" {
		return false
	}

	return true
}