package core

import (
	"gvb/site"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func InitSiteInfo(siteInfoPath string) site.SieInfo {
	file, err := ioutil.ReadFile(siteInfoPath)
	if err != nil {
		panic(err)
	}

	var siteInfo site.SieInfo
	if err := yaml.Unmarshal(file, &siteInfo); err != nil {
		panic(err)
	}

	return siteInfo
}
