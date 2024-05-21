package res

type OssConfig struct {
	Domain     string `json:"domain"`
	Bucket     string `json:"bucket"`
	Region     string `json:"region"`
	KeyPrefix  string `json:"keyprefix"`
	ImgProcess string `json:"imgprocess"`
	UpToken    string `json:"uptoken"`
}
