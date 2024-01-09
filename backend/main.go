package main

import (
	"backend/internal/core"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile = flag.String("f", "etc/dev.yaml", "the config file")

func main() {
	flag.Parse()
	core.NewServiceContext(*configFile)
	serverSvc := core.GetSvcContext()
	fmt.Println("初始化server完成")
	addr := fmt.Sprintf("%s:%d", serverSvc.Config.Host, serverSvc.Config.Port)
	// 后台异步启动gin服务
	go func() {
		panic(serverSvc.Server.Run(addr))
	}()

	// 后台协程监控系统负载情况

	// 优雅启停 使用signal来捕捉系统信号
	ctx, cancel := context.WithCancel(context.Background())
	// 监听系统信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	// 等待系统信号
	select {
	case sig := <-signalChan:
		fmt.Println("Received signal:", sig)
	case <-ctx.Done():
		fmt.Println("Context canceled")
	}

	// 停止 HTTP 服务器
	fmt.Println("正在关闭gin HTTP 服务 ...")
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("服务已经被关闭")
}
