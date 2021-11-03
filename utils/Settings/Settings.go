package Settings

import (
	"strconv"
	"strings"
	"xyz.nyan/MediaWiki-Bot/Struct"
)

func Settings(SNSName string, QQMessagejson Struct.QQWebHook_root, CommandText string) (string, bool) {
	if find := strings.Contains(CommandText, "language"); find {
		switch SNSName {
		case "QQ":
			countSplit := strings.SplitN(CommandText, " ", 2)
			Language := countSplit[1]
			UserID := QQMessagejson.Sender.Id
			return LanguageSettings(strconv.Itoa(UserID), Language)
		}
	}
	return "", false
}
