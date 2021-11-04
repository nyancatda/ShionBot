package Settings

import (
	"strings"

	"xyz.nyan/MediaWiki-Bot/src/Struct"
	"xyz.nyan/MediaWiki-Bot/src/utils"
	Languages "xyz.nyan/MediaWiki-Bot/src/utils/Language"
)

func LanguageSettings(UserID string, Language string) (string, bool) {
	var MessageOK bool
	var Message string
	db := utils.SQLLiteLink()

	db.AutoMigrate(&Struct.QQUserInfo{})

	var user Struct.QQUserInfo
	db.Where("account = ?", UserID).Find(&user)
	if user.Account != UserID {
		files, _, _ := utils.GetFilesAndDirs("./resources/language/")
		for _, dir := range files {
			LanguageName := strings.Replace(dir, `\`, "/", 1)
			LanguageNames := strings.Split(LanguageName, "/")
			LanguageNames = strings.Split(LanguageNames[4], ".")
			LanguageName = LanguageNames[0]
			if LanguageName == Language {
				UserInfos := Struct.QQUserInfo{Account: UserID, Language: Language}
				db.Create(&UserInfos)
				MessageOK = true
				Message = Languages.StringVariable(1, Languages.Message(UserID).LanguageModifiedSuccessfully, LanguageName, "")
				break
			} else {
				MessageOK = true
				Message = Languages.StringVariable(1, Languages.Message(UserID).LanguageModificationFailed, LanguageName, "")
			}
		}
	} else {
		files, _, _ := utils.GetFilesAndDirs("./resources/language/")
		for _, dir := range files {
			LanguageName := strings.Replace(dir, `\`, "/", 1)
			LanguageNames := strings.Split(LanguageName, "/")
			LanguageNames = strings.Split(LanguageNames[4], ".")
			LanguageName = LanguageNames[0]
			if Language == LanguageName {
				db.Model(&Struct.QQUserInfo{}).Where("account = ?", UserID).Update("language", Language)
				MessageOK = true
				Message = Languages.StringVariable(1, Languages.Message(UserID).LanguageModifiedSuccessfully, LanguageName, "")
				break
			} else {
				MessageOK = true
				Message = Languages.StringVariable(1, Languages.Message(UserID).LanguageModificationFailed, LanguageName, "")
			}
		}
	}

	return Message, MessageOK
}
