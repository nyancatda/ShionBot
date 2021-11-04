package Struct

type WebHookJson struct {
	Type         string        `json:"type"`
	Sender       Sender        `json:"sender"`
	FromId       int           `json:"fromId"`
	Target       int           `json:"target"`
	MessageChain []interface{} `json:"messageChain"`
	Subject      Subject       `json:"subject"`
}
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
