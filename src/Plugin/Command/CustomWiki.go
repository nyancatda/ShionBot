package Command

import (
	"encoding/json"
	"strings"

	"xyz.nyan/ShionBot/src/MediaWikiAPI"
	"xyz.nyan/ShionBot/src/Struct"
	"xyz.nyan/ShionBot/src/utils"
	"xyz.nyan/ShionBot/src/utils/Language"
)

func WikiAdd(SNSName string, UserID string, CommandText string) (string, bool) {
	var MessageOK bool
	var Message string

	if find := strings.Contains(CommandText, " "); find {
		CommandParameter := strings.SplitN(CommandText, " ", 3)
		if len(CommandParameter) != 3 {
			Message = Language.Message(SNSName, UserID).WikiAddHelp
			MessageOK = true
			return Message, MessageOK
		}
		NewWikiName := CommandParameter[1]
		NewWikiLink := CommandParameter[2]

		WikiSiteinfo, err := MediaWikiAPI.QuerySiteinfoGeneral("http://" + NewWikiLink)
		if err != nil {
			Message = Language.Message(SNSName, UserID).WikiAddFailed
			MessageOK = true
			return Message, MessageOK
		}
		if _, ok := WikiSiteinfo["query"].(map[string]interface{})["general"].(map[string]interface{})["sitename"]; !ok {
			Message = Language.Message(SNSName, UserID).WikiAddFailed
			MessageOK = true
			return Message, MessageOK
		}

		db := utils.SQLLiteLink()
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
			Message = utils.StringVariable(Language.Message(SNSName, UserID).WikiAddSucceeded, []string{NewWikiName, NewWikiLink})
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
				Message = utils.StringVariable(Language.Message(SNSName, UserID).WikiAddSucceeded, []string{NewWikiName, NewWikiLink})
			} else {
				OldWikiInfoData := user.WikiInfo
				WikiInfoData := []interface{}{}
				json.Unmarshal([]byte(OldWikiInfoData), &WikiInfoData)
				//检查添加是否重复
				for _, value := range WikiInfoData {
					OldWikiName := value.(map[string]interface{})["WikiName"]
					if OldWikiName == NewWikiName {
						MessageOK = true
						Message = utils.StringVariable(Language.Message(SNSName, UserID).WikiAddRepeat, []string{NewWikiName})
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
				Message = utils.StringVariable(Language.Message(SNSName, UserID).WikiAddSucceeded, []string{NewWikiName, NewWikiLink})
			}
		}
	} else {
		if CommandText == "wikiadd" {
			Message = Language.Message(SNSName, UserID).WikiAddHelp
			MessageOK = true
		}
	}

	return Message, MessageOK
}
