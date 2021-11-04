package Settings

import (
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
		files := Languages.LanguageList()
		for _, LanguageName := range files {
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
		files := Languages.LanguageList()
		for _, LanguageName := range files {
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
