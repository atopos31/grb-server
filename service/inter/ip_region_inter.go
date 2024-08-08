package inter

// IPtoRegion
type IPtools interface {
	GetRegion(ip string) (string, error) // 获取ip地址对应的区域
}
