package system

import (
	"gvb/app"
	"gvb/models/res"
	"io"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

func getInfo(c *gin.Context) {
	hostInfo, err := host.Info()
	if err != nil {
		res.Error(c, err)
		return
	}
	info := &res.Info{
		Arch:            hostInfo.KernelArch,
		KernelVersion:   hostInfo.KernelVersion,
		OS:              hostInfo.Platform,
		Platform:        hostInfo.PlatformFamily,
		PlatformVersion: hostInfo.PlatformVersion,
	}
	res.Success(c, info)
}

type step func(w io.Writer) bool

// @Summary 服务器状态监控SSE接口
// @Tags 系统信息
// TODO
func handlerServerStatus(c *gin.Context) {
	step := statusWriter(c)
	isClose := c.Stream(step)
	if isClose {
		app.Log.Warningf("Client IP:%s Close connent\n", c.RemoteIP())
	} else {
		app.Log.Warningf("Server Close connent\n")
	}
}

func statusWriter(c *gin.Context) step {
	return func(w io.Writer) bool {
		memStatus, _ := mem.VirtualMemory()
		host := res.Cmn{
			Mem: res.Mem{
				Total:       memStatus.Total,
				Used:        memStatus.Used,
				UsedPercent: int(memStatus.UsedPercent),
			},
		}
		var wg sync.WaitGroup
		wg.Add(2)
		go loadCPUperent(&wg, &host.CpuPrecent, 2*time.Second)
		go loadNetIO(&wg, &host.Net, 2*time.Second)
		wg.Wait()
		c.SSEvent("message", host)
		if app.Exit {
			return false
		}
		return true
	}
}

func loadCPUperent(wg *sync.WaitGroup, percent *int, interval time.Duration) {
	defer wg.Done()
	totalPercent, _ := cpu.Percent(interval, true)
	*percent = int(totalPercent[0])
	return
}

func loadNetIO(wg *sync.WaitGroup, netIO *res.Net, interval time.Duration) {
	defer wg.Done()
	st, _ := net.IOCounters(false)
	lastSend, lastRecv := st[0].BytesSent, st[0].BytesRecv
	time.Sleep(interval)
	st, _ = net.IOCounters(false)
	netIO.BytesRecv = (st[0].BytesRecv - lastRecv) / uint64(interval.Seconds())
	netIO.BytesSent = (st[0].BytesSent - lastSend) / uint64(interval.Seconds())
	return
}
