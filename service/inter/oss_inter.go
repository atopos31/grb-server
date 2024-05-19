package inter

// Oss服务接口
type OssService interface {
	GetUploadToken() (string, error)
}
