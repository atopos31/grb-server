package inter

import "gvb/models/res"

// Oss服务接口
type OssService interface {
	GetUploadToken() (res.OssConfig, error) // 获取上传token
}
