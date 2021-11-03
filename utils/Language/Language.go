package Language

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
	"xyz.nyan/MediaWiki-Bot/utils"
)

func ReadLanguage(Language string) *LanguageInfo {
	content, err := ioutil.ReadFile("./language/" + Language + ".yml")
	if err != nil {
		panic(err)
	}
	newStu := &LanguageInfo{}
	err = yaml.Unmarshal(content, &newStu)
	if err != nil {
		panic(err)
	}
	return newStu
}

//替换字符串中的变量位置
func StringVariable(quantity int, strHaiCoder string, text0 string, text1 string) string {
	text := ""
	switch quantity {
	case 1:
		text = strings.Replace(strHaiCoder, "{%0}", text0, 1)
	case 2:
		text = strings.Replace(strHaiCoder, "{%0}", text0, 1)
		text = strings.Replace(text, "{%1}", text1, 1)
	}
	return text
}

//释放语言文件
func ReleaseFile() {
	//打包语言文件
	//go-bindata -o=utils/Language/languages.go -pkg=Language language/...
	_, err := os.Stat("./language")
	if err != nil {
		os.Mkdir("./language", 0777)
		for filename := range _bindata {
			bytes, _ := Asset(filename)
			ioutil.WriteFile(filename, bytes, 0664)
		}
	}
}

func Message() *LanguageInfo {
	Config := utils.ReadConfig()
	language := Config.Run.Language
	Info := ReadLanguage(language)
	return Info
}
