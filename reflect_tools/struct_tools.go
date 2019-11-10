/**
 * @Author: DollarKiller
 * @Description: structTools
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:23 2019-11-08
 */
package reflect_tools

import (
	"errors"
	"reflect"
)

// 结构体相关工具
type StructTools struct {
}

// 结构体转Map
func (s *StructTools) StructToMap(obj interface{}) (result map[string]interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			result = nil
			err = errors.New("Structure definition error")
		}
	}()

	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		// 如果是指针
		typ := reflect.TypeOf(obj).Elem()
		val := reflect.ValueOf(obj).Elem()
		if typ.Kind() != reflect.Struct {
			return nil, errors.New("this is not struct")
		}
		result = map[string]interface{}{}
		for i := 0; i < typ.NumField(); i++ {
			c := val.Field(i).Kind()
			if c == reflect.Struct || c == reflect.Ptr {
				results, err := s.StructToMap(val.Field(i).Interface())
				if err != nil {
					continue
				}
				result[typ.Field(i).Name] = results
			} else {
				result[typ.Field(i).Name] = val.Field(i).Interface()
			}
		}
		return result, nil
	} else if typ.Kind() == reflect.Struct {
		// 如果是结构体
		val := reflect.ValueOf(obj)
		if typ.Kind() != reflect.Struct {
			return nil, errors.New("this is not struct")
		}
		result = map[string]interface{}{}
		for i := 0; i < typ.NumField(); i++ {
			c := val.Field(i).Kind()
			if c == reflect.Struct || c == reflect.Ptr {
				results, err := s.StructToMap(val.Field(i).Interface())
				if err != nil {
					continue
				}
				result[typ.Field(i).Name] = results
			} else {
				result[typ.Field(i).Name] = val.Field(i).Interface()
			}
		}
		return result, nil
	}

	return nil, errors.New("this is not struct")
}
