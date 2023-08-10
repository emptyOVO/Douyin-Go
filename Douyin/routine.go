package main

import (
	"Douyin/controller"
	"Douyin/middleware"
	size "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	// 基础接口
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", middleware.TokenAuth(), controller.UserInfo)
	apiRouter.POST("/user/register/", middleware.Check(), controller.Register) //中间件检查账号密码是否合法
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", middleware.TokenAuth(), size.RequestSizeLimiter(4<<20), controller.Publish) //size中间件限制请求大小，这里为3mb
	apiRouter.GET("/publish/list/", middleware.TokenAuth(), controller.PublishList)
	//互动接口
	apiRouter.POST("/favorite/action/", middleware.TokenAuth(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.TokenAuth(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.TokenAuth(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middleware.TokenAuth(), controller.CommentList)
	//社交接口
	apiRouter.POST("/relation/action/", middleware.TokenAuth(), controller.FollowAction)
	apiRouter.GET("/relation/follow/list/", middleware.TokenAuth(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.TokenAuth(), controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", middleware.TokenAuth(), controller.FriendList)
	apiRouter.GET("/message/chat/", middleware.TokenAuth(), controller.MessageChat)
	apiRouter.POST("/message/action/", middleware.TokenAuth(), controller.MessageAction)
}
