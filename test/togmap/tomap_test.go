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
			}
		]
	}
	`

	mapun, e := gmap.Unmarshal(data)
	if e != nil {
		panic(e)
	}

	s, b := mapun.GetString("device")
	if b {
		log.Println(s)
	}

	i3, i4 := mapun.GetMap("ppc")
	if i4 {
		log.Println(i3)
	}

	getMap, i := mapun.GetSlice("data")
	if i {
		log.Println(getMap)
	}

	i2,bo := mapun.GetMap2(getMap[0])
	if bo {
		log.Println(i2)
	}
}
