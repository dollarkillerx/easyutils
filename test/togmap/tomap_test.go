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

	getMap, i := mapun.GetMap("data")
	if i {
		log.Println(getMap)
	}
}

