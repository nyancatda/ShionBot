package InformationProcessing

import (
	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

func KaiHeiLaWebHookVerifyProcessing(c *gin.Context, json Struct.WebHookJson) {
	Challenge := json.D.Challenge
	JsonData := map[string]interface{}{
		"challenge": Challenge,
	}
	c.JSONP(200, JsonData)
}
