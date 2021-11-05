package InformationProcessing

import (
	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

func InformationProcessing(json Struct.WebHookJson) {
	if json.Type != "" {
		go QQMessageProcessing(json)
		return
	}

	if json.Update_id != 0 {
		go TelegramMessageProcessing(json)
	}

	if json.Destination != "" {
		go LineMessageProcessing(json)
	}
}