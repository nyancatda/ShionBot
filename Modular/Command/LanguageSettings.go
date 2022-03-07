package Command

import (
	"strings"

	"github.com/nyancatda/ShionBot/Utils"
	Languages "github.com/nyancatda/ShionBot/Utils/Language"
	"github.com/nyancatda/ShionBot/Utils/SQLDB"
)

func LanguageSettings(SNSName string, UserID string, CommandText string) (string, bool) {
	var MessageOK bool
	var Message string

	if find := strings.Contains(CommandText, " "); find {
		countSplit := strings.SplitN(CommandText, " ", 2)
		Language := countSplit[1]
		db := SQLDB.DB
		var user SQLDB.UserInfo
		db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
		if user.Account != UserID {
			LanguageExist := Languages.LanguageExist(Language)
			if LanguageExist {
				UserInfos := SQLDB.UserInfo{SNSName: SNSName, Account: UserID, Language: Language}
				db.Create(&UserInfos)
				MessageOK = true
				Message = Utils.StringVariable(Languages.Message(SNSName, UserID).LanguageModifiedSuccessfully, []string{Language})
			} else {
				MessageOK = true
				Message = Utils.StringVariable(Languages.Message(SNSName, UserID).LanguageModificationFailed, []string{Language})
			}
		} else {
			LanguageExist := Languages.LanguageExist(Language)
			if LanguageExist {
				db.Model(&SQLDB.UserInfo{}).Where("account = ? and sns_name = ?", UserID, SNSName).Update("language", Language)
				MessageOK = true
				Message = Utils.StringVariable(Languages.Message(SNSName, UserID).LanguageModifiedSuccessfully, []string{Language})
			} else {
				MessageOK = true
				Message = Utils.StringVariable(Languages.Message(SNSName, UserID).LanguageModificationFailed, []string{Language})
			}
		}
	} else {
		if CommandText == "language" {
			LanguageList := Languages.LanguageList()

			for _, LanguageName := range LanguageList {
				Message = Message + LanguageName + "\n"
			}

			Message = Utils.StringVariable(Languages.Message(SNSName, UserID).LanguageList, []string{Message})

			MessageOK = true
		}
	}

	return Message, MessageOK
}
