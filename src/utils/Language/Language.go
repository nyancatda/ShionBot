/*
 * @Author: NyanCatda
 * @Date: 2021-11-02 20:10:28
 * @LastEditTime: 2022-01-24 19:52:03
 * @LastEditors: NyanCatda
 * @Description: 语言文件处理
 * @FilePath: \ShionBot\src\Utils\Language\Language.go
 */
package Language

import (
	"io/ioutil"
	"strings"

	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils"
	"github.com/nyancatda/ShionBot/src/Utils/ReadConfig"
	"github.com/nyancatda/ShionBot/src/Utils/SQLDB"
	"gopkg.in/yaml.v2"
)

func ReadLanguage(Language string) *LanguageInfo {
	content, err := ioutil.ReadFile("./resources/language/" + Language + ".yml")
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

//获取本地语言列表
func LanguageList() []string {
	var LanguageList []string
	files, _, _ := Utils.GetFilesAndDirs("./resources/language/")
	for _, dir := range files {
		LanguageName := strings.Replace(dir, `\`, "/", 1)
		LanguageNames := strings.Split(LanguageName, "/")
		LanguageNames = strings.Split(LanguageNames[4], ".")
		LanguageName = LanguageNames[0]
		LanguageList = append(LanguageList, LanguageName)
	}
	return LanguageList
}

func LanguageExist(Language string) bool {
	files := LanguageList()
	var Exist bool
	for _, LanguageName := range files {
		if LanguageName == Language {
			Exist = true
			return Exist
		} else {
			Exist = false
		}
	}
	return Exist
}

//使用默认语言参数都为空即可
func Message(SNSName string, Account string) *LanguageInfo {
	var language string
	if Account == "" {
		Config := ReadConfig.GetConfig
		language = Config.Run.Language
	} else {
		db := SQLDB.DB
		var user Struct.UserInfo
		db.Where("account = ? and sns_name = ?", Account, SNSName).Find(&user)
		if user.Language != "" {
			language = user.Language
		} else {
			Config := ReadConfig.GetConfig
			language = Config.Run.Language
		}
	}
	Info := ReadLanguage(language)
	return Info
}

func DesignateLanguageMessage(Language string) *LanguageInfo {
	return ReadLanguage(Language)
}

func DefaultLanguageMessage() *LanguageInfo {
	Config := ReadConfig.GetConfig
	language := Config.Run.Language
	Info := ReadLanguage(language)
	return Info
}
