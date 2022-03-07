package ReleaseFile

import (
	"io/ioutil"
	"os"
)

//释放资源文件
func ReleaseFile() {
	//打包资源文件
	//go-bindata -o=Utils/ReleaseFile/resources.go -pkg=ReleaseFile resources/...

	//释放语言文件
	_, err := os.Stat("./resources/language/")
	if err != nil {
		os.MkdirAll("./resources/language/", 0777)
		for filename := range _bindata {
			bytes, _ := Asset(filename)
			ioutil.WriteFile(filename, bytes, 0664)
		}
	}
}
