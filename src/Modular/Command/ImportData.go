/*
 * @Author: NyanCatda
 * @Date: 2021-11-22 20:54:44
 * @LastEditTime: 2022-01-24 19:52:27
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \ShionBot\src\Modular\Command\ImportData.go
 */
package Command

import (
	"strings"

	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils"
	"github.com/nyancatda/ShionBot/src/Utils/Language"
	"github.com/nyancatda/ShionBot/src/Utils/SQLDB"
)

func ImportData(SNSName string, UserID string, CommandText string) (string, bool) {
	var MessageOK bool
	var Message string

	if find := strings.Contains(CommandText, " "); find {
		CommandParameter := strings.SplitN(CommandText, " ", 3)
		if len(CommandParameter) != 3 {
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/importdata", "#importdata"})
			MessageOK = true
			return Message, MessageOK
		}
		ImportSNS := CommandParameter[1]
		ImportUserID := CommandParameter[2]

		db := SQLDB.DB

		var ImportSource Struct.UserInfo
		db.Where("account = ? and sns_name = ?", ImportUserID, ImportSNS).Find(&ImportSource)
		var ImportUserInfos Struct.UserInfo
		if ImportSource.Account != ImportUserID {
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).ImportDataNull, []string{ImportUserID, ImportSNS})
			MessageOK = true
			return Message, MessageOK
		} else {
			ImportUserInfos = Struct.UserInfo{SNSName: SNSName, Account: UserID, Language: ImportSource.Language, WikiInfo: ImportSource.WikiInfo}
		}

		var user Struct.UserInfo
		db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
		if user.Account != UserID {
			db.Create(&ImportUserInfos)
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).ImportDataSucceeded, []string{ImportUserID})
			MessageOK = true
		} else {
			db.Model(&Struct.UserInfo{}).Where("account = ? and sns_name = ?", UserID, SNSName).Updates(ImportUserInfos)
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).ImportDataSucceeded, []string{ImportUserID})
			MessageOK = true
		}
	} else {
		if CommandText == "importdata" {
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/importdata", "#importdata"})
			MessageOK = true
		}
	}

	return Message, MessageOK
}
