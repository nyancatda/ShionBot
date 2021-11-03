package Settings

import (
	"strings"
	"xyz.nyan/MediaWiki-Bot/Struct"
)

func Settings(Messagejson Struct.QQWebHook_root, CommandText string) (string, bool) {
	if find := strings.Contains(CommandText, "language"); find {
		countSplit := strings.SplitN(CommandText, " ", 2)
		Language := countSplit[1]
		return LanguageSettings(Messagejson, Language)
	}
	return "", false
}
