/*
 * @Author: NyanCatda
 * @Date: 2021-11-03 15:59:39
 * @LastEditTime: 2022-01-24 19:33:48
 * @LastEditors: NyanCatda
 * @Description: 文件处理工具
 * @FilePath: \ShionBot\src\Utils\FilesAndDirs.go
 */
package Utils

import (
	"io/ioutil"
	"os"
)

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, dirs, nil
}
