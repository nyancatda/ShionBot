/*
 * @Author: NyanCatda
 * @Date: 2021-11-17 15:37:34
 * @LastEditTime: 2022-01-24 19:52:17
 * @LastEditors: NyanCatda
 * @Description: Wiki命令操作
 * @FilePath: \ShionBot\src\Modular\Command\CustomWiki.go
 */
package Command

import (
	"encoding/json"
	"strings"

	"github.com/nyancatda/ShionBot/src/MediaWikiAPI"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils"
	"github.com/nyancatda/ShionBot/src/Utils/Language"
	"github.com/nyancatda/ShionBot/src/Utils/SQLDB"
)

func WikiAdd(SNSName string, UserID string, CommandText string) (string, bool) {
	var MessageOK bool
	var Message string

	if find := strings.Contains(CommandText, " "); find {
		CommandParameter := strings.SplitN(CommandText, " ", 3)
		if len(CommandParameter) != 3 {
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/wikiadd", "#wikiadd"})
			MessageOK = true
			return Message, MessageOK
		}
		NewWikiName := CommandParameter[1]
		NewWikiLink := CommandParameter[2]

		WikiSiteinfo, err := MediaWikiAPI.QuerySiteinfoGeneral("https://" + NewWikiLink)
		if err != nil {
			Message = Language.Message(SNSName, UserID).WikiAddFailed
			MessageOK = true
			return Message, MessageOK
		}
		if WikiSiteinfo.Query.General.Sitename == "" {
			Message = Language.Message(SNSName, UserID).WikiAddFailed
			MessageOK = true
			return Message, MessageOK
		}

		db := SQLDB.DB
		var user Struct.UserInfo
		db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
		if user.Account != UserID {
			WikiInfoData := make([]map[string]string, 1)
			WikiInfoData[0] = map[string]string{
				"WikiName": NewWikiName,
				"WikiLink": NewWikiLink,
			}
			WikiInfo, _ := json.Marshal(WikiInfoData)
			UserInfos := Struct.UserInfo{SNSName: SNSName, Account: UserID, WikiInfo: string(WikiInfo)}
			db.Create(&UserInfos)
			MessageOK = true
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiAddSucceeded, []string{NewWikiName, NewWikiLink})
		} else {
			if user.WikiInfo == "" {
				WikiInfoData := make([]map[string]string, 1)
				WikiInfoData[0] = map[string]string{
					"WikiName": NewWikiName,
					"WikiLink": NewWikiLink,
				}
				WikiInfo, _ := json.Marshal(WikiInfoData)
				db.Model(&Struct.UserInfo{}).Where("account = ? and sns_name = ?", UserID, SNSName).Update("wiki_info", string(WikiInfo))
				MessageOK = true
				Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiAddSucceeded, []string{NewWikiName, NewWikiLink})
			} else {
				OldWikiInfoData := user.WikiInfo
				WikiInfoData := []interface{}{}
				json.Unmarshal([]byte(OldWikiInfoData), &WikiInfoData)
				//检查添加是否重复
				for _, value := range WikiInfoData {
					OldWikiName := value.(map[string]interface{})["WikiName"]
					if OldWikiName == NewWikiName {
						MessageOK = true
						Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiAddRepeat, []string{NewWikiName})
						return Message, MessageOK
					}
				}
				NewWikiInfoData := map[string]string{
					"WikiName": NewWikiName,
					"WikiLink": NewWikiLink,
				}
				WikiInfoData = append(WikiInfoData, NewWikiInfoData)
				WikiInfo, _ := json.Marshal(WikiInfoData)
				db.Model(&Struct.UserInfo{}).Where("account = ? and sns_name = ?", UserID, SNSName).Update("wiki_info", string(WikiInfo))
				MessageOK = true
				Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiAddSucceeded, []string{NewWikiName, NewWikiLink})
			}
		}
	} else {
		if CommandText == "wikiadd" {
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/wikiadd", "#wikiadd"})
			MessageOK = true
		}
	}

	return Message, MessageOK
}

func WikiUpdate(SNSName string, UserID string, CommandText string) (string, bool) {
	var MessageOK bool
	var Message string

	if find := strings.Contains(CommandText, " "); find {
		CommandParameter := strings.SplitN(CommandText, " ", 3)
		if len(CommandParameter) != 3 {
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/wikiupdate", "#wikiupdate"})
			MessageOK = true
			return Message, MessageOK
		}
		NewWikiName := CommandParameter[1]
		NewWikiLink := CommandParameter[2]

		WikiSiteinfo, err := MediaWikiAPI.QuerySiteinfoGeneral("https://" + NewWikiLink)
		if err != nil {
			Message = Language.Message(SNSName, UserID).WikiUpdateFailed
			MessageOK = true
			return Message, MessageOK
		}
		if WikiSiteinfo.Query.General.Sitename == "" {
			Message = Language.Message(SNSName, UserID).WikiAddFailed
			MessageOK = true
			return Message, MessageOK
		}

		db := SQLDB.DB
		var user Struct.UserInfo
		db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
		if user.Account != UserID {
			MessageOK = true
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiUpdateFailedNothingness, []string{NewWikiName})
		} else {
			if user.WikiInfo == "" {
				MessageOK = true
				Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiUpdateFailedNothingness, []string{NewWikiName})
			} else {
				OldWikiInfoData := user.WikiInfo
				WikiInfoData := []interface{}{}
				json.Unmarshal([]byte(OldWikiInfoData), &WikiInfoData)
				//检查是否存在
				i := 0
				Existence := false
				for _, value := range WikiInfoData {
					OldWikiName := value.(map[string]interface{})["WikiName"]
					if OldWikiName == NewWikiName {
						WikiInfoData[i] = map[string]string{
							"WikiName": NewWikiName,
							"WikiLink": NewWikiLink,
						}
						Existence = true
					}
					i = i + 1
				}
				if !Existence {
					MessageOK = true
					Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiUpdateFailedNothingness, []string{NewWikiName})
					return Message, MessageOK
				}
				WikiInfo, _ := json.Marshal(WikiInfoData)
				db.Model(&Struct.UserInfo{}).Where("account = ? and sns_name = ?", UserID, SNSName).Update("wiki_info", string(WikiInfo))
				MessageOK = true
				Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiUpdateSucceeded, []string{NewWikiName, NewWikiLink})
			}
		}
	} else {
		if CommandText == "wikiupdate" {
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/wikiupdate", "#wikiupdate"})
			MessageOK = true
		}
	}

	return Message, MessageOK
}

func WikiDelete(SNSName string, UserID string, CommandText string) (string, bool) {
	var MessageOK bool
	var Message string

	if find := strings.Contains(CommandText, " "); find {
		CommandParameter := strings.SplitN(CommandText, " ", 3)
		if len(CommandParameter) != 2 {
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/wikidelete", "#wikidelete"})
			MessageOK = true
			return Message, MessageOK
		}
		NewWikiName := CommandParameter[1]

		db := SQLDB.DB
		var user Struct.UserInfo
		db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
		if user.Account != UserID {
			MessageOK = true
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiDeleteFailedNothingness, []string{NewWikiName})
		} else {
			if user.WikiInfo == "" {
				MessageOK = true
				Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiDeleteFailedNothingness, []string{NewWikiName})
			} else {
				OldWikiInfoData := user.WikiInfo
				WikiInfoData := []interface{}{}
				json.Unmarshal([]byte(OldWikiInfoData), &WikiInfoData)
				//检查是否存在
				i := 0
				Existence := false
				for _, value := range WikiInfoData {
					OldWikiName := value.(map[string]interface{})["WikiName"]
					if OldWikiName == NewWikiName {
						WikiInfoData = append(WikiInfoData[:0], WikiInfoData[0+1:]...)
						Existence = true
					}
					i = i + 1
				}
				if !Existence {
					MessageOK = true
					Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiDeleteFailedNothingness, []string{NewWikiName})
					return Message, MessageOK
				}
				WikiInfo, _ := json.Marshal(WikiInfoData)
				db.Model(&Struct.UserInfo{}).Where("account = ? and sns_name = ?", UserID, SNSName).Update("wiki_info", string(WikiInfo))
				MessageOK = true
				Message = Utils.StringVariable(Language.Message(SNSName, UserID).WikiDeleteSucceeded, []string{NewWikiName})
			}
		}
	} else {
		if CommandText == "wikidelete" {
			Message = Utils.StringVariable(Language.Message(SNSName, UserID).CommandHelp, []string{"/wikidelete", "#wikidelete"})
			MessageOK = true
		}
	}

	return Message, MessageOK
}
