package Struct

type QQWebHook_root struct {
	Type         string        `json:"type"`
	Sender       SenderJson    `json:"sender"`
	FromId       int           `json:"fromId"`
	Target       int           `json:"target"`
	MessageChain []interface{} `json:"messageChain"`
	Subject      SubjectJson   `json:"subject"`
}
type SenderJson struct {
	Id                 int       `json:"id"`
	MemberName         string    `json:"memberName"`
	SpecialTitle       string    `json:"specialTitle"`
	Permission         string    `json:"permission"`
	JoinTimestamp      int       `json:"joinTimestamp"`
	LastSpeakTimestamp int       `json:"lastSpeakTimestamp"`
	MuteTimeRemaining  int       `json:"muteTimeRemaining"`
	Nickname           string    `json:"nickname"`
	Remark             string    `json:"remark"`
	Group              GroupJson `json:"group"`
}
type GroupJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type SubjectJson struct {
	Id   int    `json:"id"`
	Kind string `json:"kind"`
}