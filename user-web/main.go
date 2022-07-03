package main

import (
	"fmt"
	"github.com/ShiyuCheng2018/mxshop-api/user-web/initializer"
	"go.uber.org/zap"
)

func main() {
	port := 8021

	// initialize Logger
	initializer.InitLogger()
	// initialize Routers
	Router := initializer.Routers()

	zap.S().Infof("[Runing Server]: port: %d", port)

	err := Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Panic("[Failed Server]: ", err.Error())
	}
}
