package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	ListenAddr  = flag.String("listen", "0.0.0.0:9999", "server listen address")
	ContextPath = flag.String("contextPath", "", "the web server context path")
	ProjectName = flag.String("projectName", "ginDemo", "the web server project name")
	LogFile     = flag.String("logFile", "./ginDemo.log", "the web service log file")
	Help        = flag.Bool("help", false, "show context-sensitive help info")
	Version     = flag.Bool("version", false, "show server version info")
)

var buildVersion = "1.0"

func main() {

	flag.Parse()

	if *Help {
		flag.Usage()
		return
	}

	if *Version {
		fmt.Printf("gin demo buildVersion=%s\r\n", buildVersion)
		return
	}

	//创建日志文件
	f, _ := os.Create(*LogFile)
	//如果需要同时将日志写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	contextPath := *ContextPath

	//创建一个默认的路由引擎
	r := gin.Default()
	r.GET(contextPath+"/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello  world",
		})
	})

	r.GET(contextPath+"/health", func(c *gin.Context) {
		timeNow := time.Now()
		timeFormat := timeNow.Format("2006-01-02 15:04:05")

		result := *ProjectName + " is ok, current time is " + timeFormat
		c.String(http.StatusOK, result)
	})

	r.Run(*ListenAddr)
}
