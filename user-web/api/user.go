package api

import (
	"context"
	"fmt"
	"github.com/ShiyuCheng2018/mxshop-api/user-web/forms"
	"github.com/ShiyuCheng2018/mxshop-api/user-web/global"
	"github.com/ShiyuCheng2018/mxshop-api/user-web/global/responses"
	"github.com/ShiyuCheng2018/mxshop-api/user-web/proto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttpError(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "[Internal Server Error]",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "[Invalid Argument]",
				})
			case codes.Unavailable:
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"msg": "[Service Unavailable]",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Message(),
				})
			}
			return
		}
	}
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func GetUserList(ctx *gin.Context) {
	// connect to client grpc server
	userConnection, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList]: Failed to connect")
		HandleGrpcErrorToHttpError(err, ctx)
	}

	// get grpc client
	userSrvClient := proto.NewUserClient(userConnection)

	page := ctx.DefaultQuery("page", "0")
	pageInt, _ := strconv.Atoi(page)
	pageSize := ctx.DefaultQuery("size", "10")
	pageSizeInt, _ := strconv.Atoi(pageSize)

	response, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Page: uint32(pageInt),
		Size: uint32(pageSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]: Failed to get user list")
		HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, val := range response.Data {
		user := responses.UserInfoResponse{
			Id:       val.Id,
			Name:     val.Name,
			Gender:   val.Gender,
			Birthday: responses.JsonTime(time.Unix(int64(val.Birthday), 0)),
			Nickname: val.Nickname,
			Mobile:   val.Mobile,
			Role:     val.Role,
		}

		result = append(result, user)
	}

	ctx.JSON(http.StatusOK, result)
}

func PasswordLogin(c *gin.Context) {
	// form validator
	passwordLoginForm := forms.PasswordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorError(c, err)
	}

}
