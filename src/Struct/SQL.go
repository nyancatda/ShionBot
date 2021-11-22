package Struct

type UserInfo struct {
	Id       int    `gorm:"primary_key"`
	SNSName  string `gorm:"type:varchar(255);not null;index:sns_name"`
	Account  string `gorm:"type:varchar(255);not null;index:account"`
	Language string `gorm:"language;type:varchar(255)"`
	WikiInfo string `gorm:"wikiinfo"`
}
