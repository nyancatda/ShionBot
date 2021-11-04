package utils

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