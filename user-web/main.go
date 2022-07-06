package main

import (
	"fmt"
	"github.com/ShiyuCheng2018/mxshop-api/user-web/global"
	"github.com/ShiyuCheng2018/mxshop-api/user-web/initializer"
	"go.uber.org/zap"
)

func main() {

	// initialize Logger
	initializer.InitLogger()

	// initialize global config
	initializer.InitConfig()

	// initialize Routers
	Router := initializer.Routers()

	// initializer translator
	if err := initializer.InitTrans("en"); err != nil {
		panic(err)
	}

	zap.S().Infof("[Runing Server]: port: %d", global.ServerConfig.Port)

	err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
	if err != nil {
		zap.S().Panic("[Failed Server]: ", err.Error())
	}
}
