/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:55 2019-10-08
 */
package togmap

import (
	"github.com/dollarkillerx/easyutils/gmap"
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	data := `
	{
		"device": "this is device",
		"ppc": {
			"ok":"ppc"
		},	
		"data": [
			{
				"humidity": "this is humidity",
				"time": "this is time"
			},
			"hello"
		]
	}
	`

	mapun, e := gmap.Unmarshal(data)
	if e != nil {
		panic(e)
	}

	// 获取string
	s, b := mapun.GetString("device")
	if b {
		log.Println(s)
	}

	// 获取map
	i3, i4 := mapun.GetMap("ppc")
	if i4 {
		log.Println(i3)
	}

	// 获取slice
	getMap, i := mapun.GetSlice("data")
	if i {
		log.Println(getMap)
	}

	//i2,bo := mapun.GetMap2(getMap[0])
	//if bo {
	//	log.Println(i2)
	//}

	// 获取 slice map
	sliceMap, i5 := mapun.GetSliceMap("data")
	if i5 {
		log.Println(sliceMap)
	}
}
