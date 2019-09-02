package easyutils

// 数组
func Prepend(sc *[]interface{}, value interface{}) *[]interface{} {
	result := make([]interface{}, 1)
	result[0] = value
	for _, v := range *sc {
		result = append(result, v)
	}
	return &result
}

// 数组删除
func SliceDel(sc *[]interface{}, index int) *[]interface{} {

	return sc
}
