package Command

import (
	"xyz.nyan/MediaWiki-Bot/src/Struct"
	"xyz.nyan/MediaWiki-Bot/src/utils"
	Languages "xyz.nyan/MediaWiki-Bot/src/utils/Language"
)

func LanguageSettings(SNSName string, UserID string, Language string) (string, bool) {
	var MessageOK bool
	var Message string
	db := utils.SQLLiteLink()

	db.AutoMigrate(&Struct.UserInfo{})

	var user Struct.UserInfo
	db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
	if user.Account != UserID {
		files := Languages.LanguageList()
		for _, LanguageName := range files {
			if LanguageName == Language {
				UserInfos := Struct.UserInfo{SNSName: SNSName, Account: UserID, Language: Language}
				db.Create(&UserInfos)
				MessageOK = true
				Message = Languages.StringVariable(1, Languages.Message(SNSName, UserID).LanguageModifiedSuccessfully, LanguageName, "")
				break
			} else {
				MessageOK = true
				Message = Languages.StringVariable(1, Languages.Message(SNSName, UserID).LanguageModificationFailed, LanguageName, "")
			}
		}
	} else {
		files := Languages.LanguageList()
		for _, LanguageName := range files {
			if Language == LanguageName {
				db.Model(&Struct.UserInfo{}).Where("account = ? and sns_name = ?", UserID, SNSName).Update("language", Language)
				MessageOK = true
				Message = Languages.StringVariable(1, Languages.Message(SNSName, UserID).LanguageModifiedSuccessfully, LanguageName, "")
				break
			} else {
				MessageOK = true
				Message = Languages.StringVariable(1, Languages.Message(SNSName, UserID).LanguageModificationFailed, LanguageName, "")
			}
		}
	}

	return Message, MessageOK
}
