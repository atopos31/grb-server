package site

type Social struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type User struct {
	Name   string `yaml:"name"`
	Avatar string `yaml:"avatar"`
}

type Site struct {
	Record string `yaml:"record"`
}

type SieInfo struct {
	Social []Social `yaml:"social"`
	User   User     `yaml:"user"`
	Site   Site     `yaml:"site"`
}
