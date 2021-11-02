package Language

type LanguageInfo struct {
	MainErrorTips       string `yaml:"MainErrorTips"`
	UnableApplySession  string `yaml:"UnableApplySession"`
	ConfigFileException string `yaml:"ConfigFileException"`
	CannotConnectMirai  string `yaml:"CannotConnectMirai"`
	RunOK               string `yaml:"RunOK"`
	RunOK_Port          string `yaml:"RunOK_Port"`
	Nudge               string `yaml:"Nudge"`
	WikiLinkError_1     string `yaml:"WikiLinkError_1"`
	WikiLinkError_2     string `yaml:"WikiLinkError_2"`
	HelpText            string `yaml:"HelpText"`
	GetWikiInfoError_1  string `yaml:"GetWikiInfoError_1"`
	GetWikiInfoError_2  string `yaml:"GetWikiInfoError_2"`
	GetWikiInfoError_3  string `yaml:"GetWikiInfoError_3"`
	WikiInfoSearch_1    string `yaml:"WikiInfoSearch_1"`
	WikiInfoSearch_2    string `yaml:"WikiInfoSearch_2"`
	WikiInfoSearch_3    string `yaml:"WikiInfoSearch_3"`
	WikiInfoRedirect_1  string `yaml:"WikiInfoRedirect_1"`
	WikiInfoRedirect_2  string `yaml:"WikiInfoRedirect_2"`
	WikiInfoRedirect_3  string `yaml:"WikiInfoRedirect_3"`
}
