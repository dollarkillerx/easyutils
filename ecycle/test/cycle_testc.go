/**
 * @Author: DollarKillerX
 * @Description: cycle_testc.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午1:34 2019/12/16
 */
package main

import (
	"github.com/dollarkillerx/easyutils/ecycle"
	"log"
	"time"
)

func main() {
	go ecycle.Cycle()
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case c := <-ticker.C:
			//panic(c)
			log.Println(c)
			return
		}
	}
}
