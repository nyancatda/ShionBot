package InformationProcessing

import (
	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

func InformationProcessing(json Struct.WebHookJson) {
	if json.Type != "" {
		QQMessageProcessing(json)
		return
	}

	if json.Update_id != 0 {
		TelegramMessageProcessing(json)
	}
}