package HttpAPI

import (
	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/src/HttpAPI/API"
)

var APIName string

func HttpAPIStart(r *gin.Engine) {
	r.GET("/api/:api_name", func(c *gin.Context) {
		APIName = c.Param("api_name")
		var data map[string]interface{}
		if APIName == "query" {
			data = API.QueryInfo(c)
		}
		c.JSONP(200, data)
	})
}
