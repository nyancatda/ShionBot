package InformationProcessing

import (
	"fmt"

	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

func TelegramMessageProcessing(json Struct.WebHookJson) {
	fmt.Println(json.Message.Text)
}