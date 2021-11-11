package Struct

type WebHookJson struct {
	//QQ
	Type         string        `json:"type"`
	Sender       Sender        `json:"sender"`
	FromId       int           `json:"fromId"`
	Target       int           `json:"target"`
	MessageChain []interface{} `json:"messageChain"`
	Subject      Subject       `json:"subject"`

	//Telegram
	Update_id int     `json:"update_id"`
	Message   Message `json:"message"`

	//Line
	Destination string   `json:"destination"`
	Events      []Events `json:"events"`

	//KaiHeila
	S  int `json:"s"`
	D  D   `json:"d"`
	Sn int `json:"sn"`
}

//QQ
type Sender struct {
	Id                 int    `json:"id"`
	MemberName         string `json:"memberName"`
	SpecialTitle       string `json:"specialTitle"`
	Permission         string `json:"permission"`
	JoinTimestamp      int    `json:"joinTimestamp"`
	LastSpeakTimestamp int    `json:"lastSpeakTimestamp"`
	MuteTimeRemaining  int    `json:"muteTimeRemaining"`
	Nickname           string `json:"nickname"`
	Remark             string `json:"remark"`
	Group              Group  `json:"group"`
}
type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Subject struct {
	Id   int    `json:"id"`
	Kind string `json:"kind"`
}

//Telegram
type Message struct {
	Message_id int    `json:"message_id"`
	From       From   `json:"from"`
	Chat       Chat   `json:"chat"`
	Date       int    `json:"date"`
	Text       string `json:"text"`
}
type From struct {
	Id            int    `json:"id"`
	Is_bot        bool   `json:"is_bot"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Username      string `json:"username"`
	Language_code string `json:"language_code"`
}
type Chat struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Type       string `json:"type"`
}

//Line
type Events struct {
	Type       string      `json:"type"`
	Message    MessageLine `json:"message"`
	Timestamp  int         `json:"timestamp"`
	Source     Source      `json:"source"`
	ReplyToken string      `json:"replyToken"`
	Mode       string      `json:"mode"`
}
type MessageLine struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Text string `json:"text"`
}
type Source struct {
	Type    string `json:"type"`
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

//KaiHeila WebHookVerify
type D struct {
	Type          int    `json:"type"`
	Channel_type  string `json:"channel_type"`
	Target_id     string `json:"target_id"`
	Author_id     string `json:"author_id"`
	Content       string `json:"content"`
	Extra         Extra  `json:"extra"`
	Msg_id        string `json:"msg_id"`
	Msg_timestamp int    `json:"msg_timestamp"`
	Nonce         string `json:"nonce"`
	From_type     int    `json:"from_type"`
	Challenge     string `json:"challenge"`
	Verify_token  string `json:"verify_token"`
}
type Extra struct {
	Type             int           `json:"type"`
	Code             string        `json:"code"`
	Guild_id         string        `json:"guild_id"`
	Channel_name     string        `json:"channel_name"`
	Mention          []interface{} `json:"mention"`
	Mention_all      bool          `json:"mention_all"`
	Mention_roles    []interface{} `json:"mention_roles"`
	Mention_here     bool          `json:"mention_here"`
	Author           Author        `json:"author"`
	Nonce            string        `json:"nonce"`
	Last_msg_content string        `json:"last_msg_content"`
}
type Author struct {
	Id           string        `json:"id"`
	Username     string        `json:"username"`
	Identify_num string        `json:"identify_num"`
	Online       bool          `json:"online"`
	Os           string        `json:"os"`
	Status       int           `json:"status"`
	Avatar       string        `json:"avatar"`
	Vip_avatar   string        `json:"vip_avatar"`
	Banner       string        `json:"banner"`
	Nickname     string        `json:"nickname"`
	Roles        []interface{} `json:"roles"`
	Is_vip       bool          `json:"is_vip"`
	Bot          bool          `json:"bot"`
}
