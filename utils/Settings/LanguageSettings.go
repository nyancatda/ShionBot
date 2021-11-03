package Settings

import (
	"strconv"

	"strings"
	"xyz.nyan/MediaWiki-Bot/Struct"
	"xyz.nyan/MediaWiki-Bot/utils"
)

func LanguageSettings(Messagejson Struct.QQWebHook_root, Language string) (string, bool) {
	var MessageOK bool
	var Message string
	db := utils.SQLLiteLink()

	db.AutoMigrate(&Struct.UserInfo{})

	UserID := Messagejson.Sender.Id
	var user Struct.UserInfo
	db.Where("account = ?", UserID).Find(&user)
	if user.Account != strconv.Itoa(UserID) {
		files, _, _ := utils.GetFilesAndDirs("./language")
		for _, dir := range files {
			LanguageName := strings.Replace(dir, `\`, "/", 1)
			LanguageNames := strings.Split(LanguageName, "/")
			LanguageNames = strings.Split(LanguageNames[2], ".")
			LanguageName = LanguageNames[0]
			if LanguageName == Language {
				UserInfos := Struct.UserInfo{Account: strconv.Itoa(UserID), Language: Language}
				db.Create(&UserInfos)
				MessageOK = true
				Message = "语言已更改为" + LanguageName
				break
			} else {
				MessageOK = false
				Message = LanguageName + "语言不存在"
			}
		}
	} else {
		files, _, _ := utils.GetFilesAndDirs("./language")
		for _, dir := range files {
			LanguageName := strings.Replace(dir, `\`, "/", 1)
			LanguageNames := strings.Split(LanguageName, "/")
			LanguageNames = strings.Split(LanguageNames[2], ".")
			LanguageName = LanguageNames[0]
			if Language == LanguageName {
				db.Model(&Struct.UserInfo{}).Where("account = ?", UserID).Update("language", Language)
				MessageOK = true
				Message = "语言已更改为" + LanguageName
				break
			} else {
				MessageOK = false
				Message = LanguageName + "语言不存在"
			}
		}
	}

	return Message, MessageOK
}
