package res

type Info struct {
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platform_version"`
	KernelVersion   string `json:"kernel_version"`
	Arch            string `json:"arch"`
}

type Cmn struct {
	CpuPrecent int `json:"cpuPrecent"`
	Mem        Mem `json:"mem"`
	Net        Net `json:"net"`
}

type Mem struct {
	Total       uint64 `json:"total"` // 单位 b 下同
	Used        uint64 `json:"used"`
	UsedPercent int    `json:"usedPercent"`
}

type Net struct {
	BytesRecv uint64 `json:"bytesRecv"`
	BytesSent uint64 `json:"bytesSent"`
}
