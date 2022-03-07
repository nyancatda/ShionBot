/*
 * @Author: NyanCatda
 * @Date: 2022-01-24 21:37:39
 * @LastEditTime: 2022-01-24 21:42:09
 * @LastEditors: NyanCatda
 * @Description: 数据库模型
 * @FilePath: \ShionBot\src\Utils\SQLDB\Models.go
 */
package SQLDB

type UserInfo struct {
	Id       int    `gorm:"id;type:int;primaryKey;not null;comment:ID"`
	SNSName  string `gorm:"sns_name;type:string;not null;comment:所属聊天软件"`
	Account  string `gorm:"account;type:string;not null;comment:账号"`
	Language string `gorm:"language;type:string;comment:使用语言"`
	WikiInfo string `gorm:"wikiinfo;type:string;comment:自定义Wiki信息"`
}
