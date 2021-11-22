package Command

import (
	"strings"

	"xyz.nyan/ShionBot/src/Struct"
	"xyz.nyan/ShionBot/src/utils"
	"xyz.nyan/ShionBot/src/utils/Language"
)

func ImportData(SNSName string, UserID string, CommandText string) (string, bool) {
	var MessageOK bool
	var Message string

	if find := strings.Contains(CommandText, " "); find {
		CommandParameter := strings.SplitN(CommandText, " ", 3)
		if len(CommandParameter) != 3 {
			Message = utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/importdata", "#importdata"})
			MessageOK = true
			return Message, MessageOK
		}
		ImportSNS := CommandParameter[1]
		ImportUserID := CommandParameter[2]

		db := utils.SQLLiteLink()

		var ImportSource Struct.UserInfo
		db.Where("account = ? and sns_name = ?", ImportUserID, ImportSNS).Find(&ImportSource)
		var ImportUserInfos Struct.UserInfo
		if ImportSource.Account != ImportUserID {
			Message = "不存在ID为 " + ImportUserID + " 的 " + ImportSNS + " 用户，请检查输入是否正确"
			MessageOK = true
			return Message, MessageOK
		} else {
			ImportUserInfos = Struct.UserInfo{SNSName: SNSName, Account: UserID, Language: ImportSource.Language, WikiInfo: ImportSource.WikiInfo}
		}

		var user Struct.UserInfo
		db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
		if user.Account != UserID {
			db.Create(&ImportUserInfos)
			Message = "已成功将 " + ImportUserID + " 的数据导入你的账户"
			MessageOK = true
		} else {
			db.Model(&Struct.UserInfo{}).Where("account = ? and sns_name = ?", UserID, SNSName).Updates(ImportUserInfos)
			Message = "已成功将 " + ImportUserID + " 的数据导入你的账户"
			MessageOK = true
		}
	} else {
		if CommandText == "importdata" {
			Message = utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/importdata", "#importdata"})
			MessageOK = true
		}
	}

	return Message, MessageOK
}
