package Language

type LanguageInfo struct {
	LanguageModifiedSuccessfully string `yaml:"LanguageModifiedSuccessfully"`
	LanguageModificationFailed   string `yaml:"LanguageModificationFailed"`
	MainErrorTips                string `yaml:"MainErrorTips"`
	UnableApplySession           string `yaml:"UnableApplySession"`
	ConfigFileException          string `yaml:"ConfigFileException"`
	CannotConnectMirai           string `yaml:"CannotConnectMirai"`
	RunOK                        string `yaml:"RunOK"`
	Nudge                        string `yaml:"Nudge"`
	WikiLinkError                string `yaml:"WikiLinkError"`
	HelpText                     string `yaml:"HelpText"`
	GetWikiInfoError             string `yaml:"GetWikiInfoError"`
	WikiInfoSearch               string `yaml:"WikiInfoSearch"`
	WikiInfoRedirect             string `yaml:"WikiInfoRedirect"`
	TitleEmpty                   string `yaml:"TitleEmpty"`
	LanguageList                 string `yaml:"LanguageList"`
}
