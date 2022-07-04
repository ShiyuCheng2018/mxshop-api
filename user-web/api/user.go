package api

import (
	"context"
	"fmt"
	"github.com/ShiyuCheng2018/mxshop-api/user-web/global/responses"
	"github.com/ShiyuCheng2018/mxshop-api/user-web/proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

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

func GetUserList(ctx *gin.Context) {
	ip := "127.0.0.1"
	port := 50051

	// connect to client grpc server
	userConnection, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList]: Failed to connect")
		HandleGrpcErrorToHttpError(err, ctx)
	}

	// get grpc client
	userSrvClient := proto.NewUserClient(userConnection)

	response, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Page: 0,
		Size: 0,
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
