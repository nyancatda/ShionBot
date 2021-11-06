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

	var user Struct.UserInfo
	db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
	if user.Account != UserID {
		LanguageExist := Languages.LanguageExist(Language)
		if LanguageExist {
			UserInfos := Struct.UserInfo{SNSName: SNSName, Account: UserID, Language: Language}
			db.Create(&UserInfos)
			MessageOK = true
			Message = Languages.StringVariable(1, Languages.Message(SNSName, UserID).LanguageModifiedSuccessfully, Language, "")
		} else {
			MessageOK = true
			Message = Languages.StringVariable(1, Languages.Message(SNSName, UserID).LanguageModificationFailed, Language, "")
		}
	} else {
		LanguageExist := Languages.LanguageExist(Language)
		if LanguageExist {
			db.Model(&Struct.UserInfo{}).Where("account = ? and sns_name = ?", UserID, SNSName).Update("language", Language)
			MessageOK = true
			Message = Languages.StringVariable(1, Languages.Message(SNSName, UserID).LanguageModifiedSuccessfully, Language, "")
		} else {
			MessageOK = true
			Message = Languages.StringVariable(1, Languages.Message(SNSName, UserID).LanguageModificationFailed, Language, "")
		}
	}

	return Message, MessageOK
}
