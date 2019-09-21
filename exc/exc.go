/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:27 2019-09-21
 */
package exc

import (
	"bytes"
	"github.com/dollarkillerx/easyutils/clog"
	"os/exec"
)

type Exc struct {

}

func (u *Exc) Exec(sh string, arg ...string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(sh, arg...)
	cmd.Stdout = &stdout // 输出
	cmd.Stderr = &stderr // 输出错误
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func (u *Exc) ReLan(sh string, arg ...string) error {
	cmd := exec.Command(sh, arg...)
	err := cmd.Start()
	if err != nil {
		clog.Println(err)
	}
	err = cmd.Wait()
	return err
}
