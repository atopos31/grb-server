package inter

import "gvb/models/res"

// Oss服务接口
type OssService interface {
	GetUploadToken() (res.OssConfig, error) // 获取前端上传所需配置信息
}