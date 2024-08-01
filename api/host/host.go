package host

import (
	"fmt"
	"gvb/models/res"
	"io"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

type step func(w io.Writer) bool

// @Tags 系统信息
// TODO
func GetHostInfo(c *gin.Context) {
	step := sseHost(c)
	isClose := c.Stream(step)
	if isClose {
		fmt.Printf("Client IP:%sClose connent\n", c.RemoteIP())
	} else {
		fmt.Println("Server close connnent")
	}
}

func sseHost(c *gin.Context) step {
	return func(w io.Writer) bool {
		memStatus, _ := mem.VirtualMemory()
		host := res.Host{
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
