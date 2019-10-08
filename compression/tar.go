/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:25 2019-09-21
 */
package compression

import "github.com/dollarkillerx/easyutils/exc"

type Tar struct {
}

func (t *Tar) Tar(path, des string) error {
	//tar zcvf FileName.tar.gz DirName
	ex := exc.Exc{}
	err := ex.ReLan("tar", "zcvf", des+".tar.gz", path)
	return err
}

func (t *Tar) UnTar(des, path string) error {
	//tar zxvf FileName.tar.gz
	ex := exc.Exc{}
	err := ex.ReLan("tar", "zxvf", des, path)
	return err
}
