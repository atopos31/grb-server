package res

type Host struct {
	CpuPrecent int `json:"cpuPrecent"`
	Mem        Mem     `json:"mem"`
	Net        Net     `json:"net"`
}

type Mem struct {
	Total       uint64  `json:"total"` // 单位 b 下同
	Used        uint64  `json:"used"`
	UsedPercent int `json:"usedPercent"`
}

type Net struct {
	BytesRecv uint64 `json:"bytesRecv"`
	BytesSent uint64 `json:"bytesSent"`
}
