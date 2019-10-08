/**
 * @Author: DollarKiller
 * @Description: gmap json to map
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:51 2019-10-08
 */
package gmap

import "encoding/json"

type mapun struct {
	dataStr string
	mc      map[string]interface{}
}

func Unmarshal(jsn string) (*mapun, error) {
	m := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsn), &m)
	if err != nil {
		return nil, err
	}

	return &mapun{
		dataStr: jsn,
		mc:      m,
	}, nil
}

func (m *mapun) GetString(key string) (string,bool) {
	s,ok := m.mc[key].(string)
	return s,ok
}

func (m *mapun) GetInt(key string) (int,bool) {
	s,ok := m.mc[key].(int)
	return s,ok
}

func (m *mapun) GetMap (key string) (map[string]interface{}, bool) {
	i,ok := m.mc[key].([]interface{})[0].(map[string]interface{})
	return i,ok
}

func (m *mapun) GetMap2 (data interface{},key string) (map[string]interface{}, bool) {
	i,ok := data.([]interface{})[0].(map[string]interface{})
	return i,ok
}


// 消耗大
//func (m *mapun) ToMap() (map[string]interface{}, bool) {
//	data := make(map[string]interface{})
//
//	for k,v := range m.mc {
//
//	}
//}