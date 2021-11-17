package Struct

type UserInfo struct {
	Id       int    `gorm:"primary_key"`
	SNSName  string `gorm:"sns_name"`
	Account  string `gorm:"type:varchar(255);not null;index:account"`
	Language string `gorm:"type:varchar(255);index:language"`
	WikiInfo string `gorm:"type:varchar(255);index:wikiinfo"`
}
