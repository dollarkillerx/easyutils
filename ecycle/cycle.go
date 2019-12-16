/**
 * @Author: DollarKillerX
 * @Description: ECycle 生命周期
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午1:27 2019/12/16
 */
package ecycle

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// cycle只能监听是否ctrl + c

func Cycle() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	for {
		select {
		case s := <-ch:
			fmt.Println(s)
			switch s {
			case syscall.SIGQUIT:
				log.Println("SIGSTOP")
				return
			case syscall.SIGSTOP:
				log.Println("SIGSTOP")
				return
			case syscall.SIGHUP:
				log.Println("SIGHUP")
				return
			case syscall.SIGKILL:
				log.Println("SIGKILL")
				return
			case syscall.SIGUSR1:
				log.Println("SIGUSR1")
				return
			default:
				log.Println("default")
				return
			}
		}

	}
}
