package res

import "gvb/site"

type BasicInfo struct {
	Social []site.Social `json:"social"`
	User   site.User     `json:"user"`
}
