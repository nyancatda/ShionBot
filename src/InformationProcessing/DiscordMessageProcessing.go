package InformationProcessing

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/bsdlp/discord-interactions-go/interactions"

	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/src/utils"
)

func DiscordWebHookConfirmationRequest(data interactions.Data) {
	response := &interactions.InteractionResponse{
		Type: interactions.ChannelMessage,
		Data: &interactions.InteractionApplicationCommandCallbackData{
			Content: "got your message kid",
		},
	}

	var responsePayload bytes.Buffer
	err := json.NewEncoder(&responsePayload).Encode(response)
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf(utils.ReadConfig().SNS.Discord.BotAPILink+"api/v8/interactions/%s/%s/callback", data.ID, data.Token)
	_, err = http.Post(url, "application/json", &responsePayload)
	if err != nil {
		fmt.Println(err)
	}
}

func DiscordAutographVerify(c *gin.Context, key ed25519.PublicKey) bool {
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

	var msg bytes.Buffer
	var body bytes.Buffer
	msg.WriteString(timestamp)

	defer func() {
		c.Request.Body = ioutil.NopCloser(&body)
	}()

	_, err = io.Copy(&msg, io.TeeReader(c.Request.Body, &body))
	if err != nil {
		return false
	}

	return ed25519.Verify(key, msg.Bytes(), sig)
}

func DiscordWebHookVerify(c *gin.Context) (bool, int, map[string]interface{}) {
	var JsonData map[string]interface{}
	hexEncodedDiscordPubkey := utils.ReadConfig().SNS.Discord.PublicKey
	discordPubkey, err := hex.DecodeString(hexEncodedDiscordPubkey)
	if err != nil {
		return false, 500, JsonData
	}

	verified := DiscordAutographVerify(c, ed25519.PublicKey(discordPubkey))
	if !verified {
		JsonData = map[string]interface{}{
			"type": 0,
		}
		return false, 401, JsonData
	}

	var data interactions.Data
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)

	buf := make([]byte, 1024)
	c.Request.Body.Read(buf)
	json.Unmarshal(bodyBytes, &data)

	c.Request.Body.Close()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if data.Token == "" {
		return false, 402, JsonData
	}

	if data.Type == interactions.Ping {
		JsonData = map[string]interface{}{
			"type": 1,
		}

		go DiscordWebHookConfirmationRequest(data)

		return true, 200, JsonData
	}
	return false, 200, JsonData
}

func DiscordWebHook(r *gin.Engine) {
	Config := utils.ReadConfig()
	WebHookKey := Config.Run.WebHookKey
	r.POST("/discord/"+WebHookKey, func(c *gin.Context) {
		//初始化WebHook验证函数
		if Bool, code, JsonData := DiscordWebHookVerify(c); Bool {
			buf := make([]byte, 1024)
			n, _ := c.Request.Body.Read(buf)
			fmt.Println(string(buf[0:n]))
			c.JSONP(code, JsonData)
		} else {
			c.JSONP(code, JsonData)
		}
	})
}
