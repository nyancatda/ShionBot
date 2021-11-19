package Command

import (
	"encoding/json"
	"strings"

	"xyz.nyan/ShionBot/src/Struct"
	"xyz.nyan/ShionBot/src/utils"
	"xyz.nyan/ShionBot/src/utils/Language"
)

func UserInfo(SNSName string, UserID string, CommandText string) (string, bool) {
	var MessageOK bool
	var Message string

	if CommandText == "userinfo" {
		db := utils.SQLLiteLink()
		var user Struct.UserInfo
		db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
		if user.Account != UserID {
			UserInfos := Struct.UserInfo{SNSName: SNSName, Account: UserID}
			db.Create(&UserInfos)
			db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
		}

		var UserWikiInfo string
		if user.WikiInfo != "" {
			WikiInfoData := user.WikiInfo
			WikiInfo := []interface{}{}
			json.Unmarshal([]byte(WikiInfoData), &WikiInfo)
			for _, value := range WikiInfo {
				UserWikiInfo = UserWikiInfo + "[[" + value.(map[string]interface{})["WikiName"].(string) + "]]" + " " + value.(map[string]interface{})["WikiLink"].(string) + "\n"
			}
			UserWikiInfo = strings.TrimRight(UserWikiInfo, "\n")
		} else {
			UserWikiInfo = Language.Message(SNSName, UserID).UserInfoNotCustomWiki
		}

		var UserLanguage string
		if user.Language != "" {
			UserLanguage = user.Language
		} else {
			Config := utils.ReadConfig()
			UserLanguage = Config.Run.Language
		}

		MessageOK = true
		Message = utils.StringVariable(Language.Message(SNSName, UserID).UserInfo, []string{user.Account, user.SNSName, UserLanguage, UserWikiInfo})
	}

	return Message, MessageOK
}
