package InformationProcessing

import (
	"strings"

	"xyz.nyan/MediaWiki-Bot/src/Struct"
	"xyz.nyan/MediaWiki-Bot/src/utils"
)

//命令处理，判断命令是否匹配，匹配则输出命令和命令参数
func CommandExtraction(SNSName string, QQjson Struct.QQWebHook_root, text string) (bool, string, string) {
	if find := strings.Contains(text, ":"); find {
		Config := utils.ReadConfig()
		var ConfigWikiName string
		for one := range Config.Wiki.([]interface{}) {
			ConfigWikiName = Config.Wiki.([]interface{})[one].(map[interface{}]interface{})["WikiName"].(string)
			if find := strings.Contains(text, ConfigWikiName); find {
				countSplit := strings.SplitN(text, ":", 2)
				Command := countSplit[0]
				Text := countSplit[1]
				return true, Text, Command
			}
		}
	} else if find := strings.Contains(text, "[["); find {
		if find := strings.Contains(text, "]]"); find {
			//获取主Wiki名字
			Config := utils.ReadConfig()
			MainWikiName := Config.Wiki.([]interface{})[0].(map[interface{}]interface{})["WikiName"].(string)

			trimStr := strings.Trim(text, "[")
			Text := strings.Trim(trimStr, "]")
			return true, Text, MainWikiName
		}
	} else if find := strings.Contains(text, "/"); find {
		switch SNSName {
		case "QQ":
			QQSettingsMessageProcessing(QQjson)
		}
		return false, "", "/"
	}

	return false, "", ""
}
