package InformationProcessing

import (
	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

func InformationProcessing(c *gin.Context, json Struct.WebHookJson) {
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

	if json.D.Challenge != "" {
		KaiHeiLaWebHookVerifyProcessing(c, json)
	}
	if json.D.Content != "" {
		go KaiHeiLaMessageProcessing(json)
	}
}
