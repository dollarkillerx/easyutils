/**
 * @Author: DollarKiller
 * @Description: gmap json to map
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:51 2019-10-08
 */
package gmap

import (
	"encoding/json"
	"errors"
)

/**
 * 重写gmap 让她更人性化
 */

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

func (m *mapun) Get(keys ...string) (interface{}, error) {
	var ic map[string]interface{}
	for nu, key := range keys {
		var e error
		var i2 interface{}
		if ic == nil {
			i2, e = m.find(m.mc, key)
			if e != nil {
				return nil, errors.New("not data")
			}
		} else {
			i2, e = m.find(ic, key)
			if e != nil {
				return nil, errors.New("not data")
			}
		}
		if nu != len(keys)-1 {
			data, ok := i2.(map[string]interface{})
			if !ok {
				return nil, errors.New("Parsing failure")
			} else {
				ic = data
				continue
			}
		} else {
			return i2, e
		}
	}
	return nil, errors.New("not data")
}

func (m *mapun) find(it map[string]interface{}, name string) (interface{}, error) {
	i, ok := it[name]
	if !ok {
		return i, errors.New("not data")
	} else {
		return i, nil
	}
}
