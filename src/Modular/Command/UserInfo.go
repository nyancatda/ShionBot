/*
 * @Author: NyanCatda
 * @Date: 2021-11-19 12:06:48
 * @LastEditTime: 2022-01-24 19:52:38
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \ShionBot\src\Modular\Command\UserInfo.go
 */
/*
 * @Author: NyanCatda
 * @Date: 2021-11-19 12:06:48
 * @LastEditTime: 2022-01-24 19:36:56
 * @LastEditors: NyanCatda
 * @Description: 用户信息命令处理
 * @FilePath: \ShionBot\src\Modular\Command\UserInfo.go
 */
package Command

import (
	"encoding/json"
	"strings"

	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils"
	"github.com/nyancatda/ShionBot/src/Utils/Language"
	"github.com/nyancatda/ShionBot/src/Utils/ReadConfig"
	"github.com/nyancatda/ShionBot/src/Utils/SQLDB"
)

func UserInfo(SNSName string, UserID string, CommandText string) (string, bool) {
	var MessageOK bool
	var Message string

	if CommandText == "userinfo" {
		db := SQLDB.DB
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
			if user.WikiInfo == "[]" {
				UserWikiInfo = Language.Message(SNSName, UserID).UserInfoNotCustomWiki
			}
		} else {
			UserWikiInfo = Language.Message(SNSName, UserID).UserInfoNotCustomWiki
		}

		var UserLanguage string
		if user.Language != "" {
			UserLanguage = user.Language
		} else {
			Config := ReadConfig.GetConfig
			UserLanguage = Config.Run.Language
		}

		MessageOK = true
		Message = Utils.StringVariable(Language.Message(SNSName, UserID).UserInfo, []string{user.Account, user.SNSName, UserLanguage, UserWikiInfo})
	}

	return Message, MessageOK
}
