package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop_api/global"
	"mxshop_api/initialize"
	"mxshop_api/utils/register/consul"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()

	// 初始化router
	Router := initialize.Routers()
	err := initialize.InitTrans("zh")
	if err != nil {
		zap.S().Panic("语言转换启动失败", err.Error())
	}

	initialize.InitSrvConn()

	viper.AutomaticEnv()
	//debug := viper.GetBool("MXSHOP_DEBUG")
	//if !debug {
	//	global.ServerConfig.Port, _ = utils.GetFreePort()
	//}

	registryClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err = registryClient.Register(global.ServerConfig.Host, global.ServerConfig.Port,
		global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册到consul失败", err.Error())
	}

	zap.S().Infof("启动服务器， 端口 %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

	//接收终止信号
	quit := make(chan os.Signal, 2)
	zap.S().Info("接收服务注销指令:", err.Error())
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = registryClient.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功:")
	}
}
