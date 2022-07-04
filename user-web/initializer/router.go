package initializer

import (
	userRouter "github.com/ShiyuCheng2018/mxshop-api/user-web/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/u/v1")
	userRouter.InitUserRouter(apiGroup)

	return router
}
