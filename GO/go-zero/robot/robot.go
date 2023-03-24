package main

import (
	"cicd.robot/robot/internal/config"
	"cicd.robot/robot/internal/handler"
	"cicd.robot/robot/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/robot-api.yaml", "the config file")

func main() {
	flag.Parse()

	// 获取工作空间根路径
	root_path, err := os.Getwd()
	if err != nil {
		log.Fatal("get root path failed")
	}
	var lc logx.LogConf
	// 从 yaml 文件中 初始化配置
	conf.MustLoad(fmt.Sprintf("%s/robot/internal/config/config.yaml", root_path), &lc)
	// logx 根据配置初始化
	logx.MustSetup(lc)
	logx.DisableStat() // 日志关闭自动收集cpu等机器信息

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
