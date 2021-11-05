package Command

import (
	"xyz.nyan/MediaWiki-Bot/src/utils/Language"
)

func Help(SNSName string, UserID string) (string, bool) {
	HelpText := Language.Message(SNSName, UserID).HelpText
	return HelpText, true
}
