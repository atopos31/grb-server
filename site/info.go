package site

type Social struct {
	Name string `yaml:"name" json:"name"`
	URL  string `yaml:"url" json:"url"`
}

type User struct {
	Name   string `yaml:"name" json:"name"`
	Avatar string `yaml:"avatar" json:"avatar"`
}

type Site struct {
	Record string `yaml:"record" json:"record"`
	Title  string `yaml:"title" json:"title"`
}

type Badge struct {
	NameLeft   string `yaml:"nameLeft" json:"nameLeft"`
	NameRight  string `yaml:"nameRight" json:"nameRight"`
	Href       string `yaml:"href" json:"href"`
	Logo       string `yaml:"logo" json:"logo"`
	ColorRight string `yaml:"colorRight" json:"colorRight"`
}

type SieInfo struct {
	Social []Social `yaml:"social" json:"social"`
	User   User     `yaml:"user" json:"user"`
	Site   Site     `yaml:"site" json:"site"`
	Badges []Badge  `yaml:"badges" json:"badges"`
}
