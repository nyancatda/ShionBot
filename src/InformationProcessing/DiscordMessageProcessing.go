package InformationProcessing

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"github.com/bsdlp/discord-interactions-go/interactions"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"net/http"
	"xyz.nyan/MediaWiki-Bot/src/utils"
)

func DiscordWebHookVerify(c *gin.Context, key ed25519.PublicKey) bool {
	var msg bytes.Buffer
	signature := c.GetHeader("X-Signature-Ed25519")
	if signature == "" {
		return false
	}
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	if len(sig) != ed25519.SignatureSize || sig[63]&224 != 0 {
		return false
	}
	timestamp := c.GetHeader("X-Signature-Timestamp")
	if timestamp == "" {
		return false
	}
	msg.WriteString(timestamp)

	var body bytes.Buffer

	defer func() {
		c.Request.Body = ioutil.NopCloser(&body)
	}()

	_, err = io.Copy(&msg, io.TeeReader(c.Request.Body, &body))
	if err != nil {
		return false
	}

	return ed25519.Verify(key, msg.Bytes(), sig)
}

func DiscordWebHook(r *gin.Engine) {
	Config := utils.ReadConfig()
	WebHookKey := Config.Run.WebHookKey
	r.POST("/discord/"+WebHookKey, func(c *gin.Context) {
		var data interactions.Data
		if err := c.ShouldBindJSON(&data); err != nil {
			return
		}
		if data.Token == "" {
			return
		}

		var JsonData map[string]interface{}
		hexEncodedDiscordPubkey := "93669706bd5f8aa9d0d6598228b83830f1c8f0a4061f4e655a89f9a9fb137b9c"
		discordPubkey, err := hex.DecodeString(hexEncodedDiscordPubkey)
		if err != nil {
			return
		}

		verified := DiscordWebHookVerify(c, ed25519.PublicKey(discordPubkey))
		if !verified {
			JsonData = map[string]interface{}{
				"type": 1,
			}
			c.JSONP(401, JsonData)
			return
		}

		if data.Type == interactions.Ping {
			JsonData = map[string]interface{}{
				"type": 1,
			}
			return
		}

		response := &interactions.InteractionResponse{
			Type: interactions.ChannelMessage,
			Data: &interactions.InteractionApplicationCommandCallbackData{
				Content: "got your message kid",
			},
		}

		var responsePayload bytes.Buffer
		err = json.NewEncoder(&responsePayload).Encode(response)
		if err != nil {
			return
		}

		_, err = http.Post(data.ResponseURL(), "application/json", &responsePayload)
		if err != nil {
			return
		}

		c.JSONP(200, JsonData)
	})
}
