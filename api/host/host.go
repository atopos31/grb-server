package host

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
)

// @Tags 系统信息
// TODO
func GetHostInfo(c *gin.Context) {
	isClose := c.Stream(func(w io.Writer) bool {
		totalPercent, err := cpu.Percent(3*time.Second, true)
		if err != nil {
			return false
		}
		c.SSEvent("message", gin.H{"time": time.Now().Format(time.RFC3339), "cpu": totalPercent})
		return true
	})
	if isClose {
		fmt.Printf("Client IP:%sClose connent\n", c.RemoteIP())
	} else {
		fmt.Println("Server close cnnnent")
	}
}
