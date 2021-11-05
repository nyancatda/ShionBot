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
