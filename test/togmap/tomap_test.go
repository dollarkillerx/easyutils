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
		"data": {
			"humidity": "this is humidity",
			"time": "this is time"
		}
	}
	`

	mapun, e := gmap.Unmarshal(data)
	if e != nil {
		panic(e)
	}
	get, e := mapun.Get("device")
	if e == nil {
		log.Println(get)
	}else {
		log.Fatalln(e)
	}

	get, e = mapun.Get("data","humidity")
	if e == nil {
		log.Println(get)
	}else {
		log.Fatalln(e)
	}
}
