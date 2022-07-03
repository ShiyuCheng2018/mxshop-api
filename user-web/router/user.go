package router

import (
	"github.com/ShiyuCheng2018/mxshop-api/user-web/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	zap.S().Info("[Router]: Configuring user router")
	{
		UserRouter.GET("list", api.GetUserList)
	}
}
