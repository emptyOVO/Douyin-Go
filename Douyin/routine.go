package main

import (
	"Douyin/controller"
	"Douyin/middleware"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", middleware.TokenAuth(), controller.UserInfo)
	//中间件检查账号密码是否合法
	apiRouter.POST("/user/register/", middleware.Check(), controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	//加入中间件限制整个post请求的大小，这里设置的是3MB Limit size of POST requests for Gin framework
	apiRouter.POST("/publish/action/", middleware.TokenAuth(), limits.RequestSizeLimiter(4<<20), controller.Publish)
	apiRouter.GET("/publish/list/", middleware.TokenAuth(), controller.PublishList)
	//  extra apis - I
	apiRouter.POST("/favorite/action/", middleware.TokenAuth(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.TokenAuth(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.TokenAuth(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middleware.TokenAuth(), controller.CommentList)
	//  social apis - II
	apiRouter.POST("/relation/action/", middleware.TokenAuth(), controller.FollowAction)
	apiRouter.GET("/relation/follow/list/", middleware.TokenAuth(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.TokenAuth(), controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", middleware.TokenAuth(), controller.FriendList)
	apiRouter.GET("/message/chat/", middleware.TokenAuth(), controller.MessageChat)
	apiRouter.POST("/message/action/", middleware.TokenAuth(), controller.MessageAction)

}
