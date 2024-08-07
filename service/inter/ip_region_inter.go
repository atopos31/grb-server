package inter

// IPtoRegion
type IPtoRegion interface {
	GetRegion(ip string) (string, error) // 获取ip地址对应的区域
}
