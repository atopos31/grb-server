package core

import (
	"gvb/site"
	"os"

	"gopkg.in/yaml.v2"
)

func InitSiteInfo(siteInfoPath string) site.SieInfo {
	file, err := os.ReadFile(siteInfoPath)
	if err != nil {
		panic(err)
	}

	var siteInfo site.SieInfo
	if err := yaml.Unmarshal(file, &siteInfo); err != nil {
		panic(err)
	}

	return siteInfo
}
