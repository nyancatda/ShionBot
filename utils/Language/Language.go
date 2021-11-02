package Language

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"xyz.nyan/MediaWiki-Bot/utils"
)

func ReadLanguage(Language string) *LanguageInfo {
	content, err := ioutil.ReadFile("./language/" + Language + ".yml")
	if err != nil {
		panic(err)
	}
	newStu := &LanguageInfo{}
	err = yaml.Unmarshal(content, &newStu)
	if err != nil {
		panic(err)
	}
	return newStu
}

func Message() *LanguageInfo {
	Config := utils.ReadConfig()
	language := Config.Run.Language
	Info := ReadLanguage(language)
	return Info
}
