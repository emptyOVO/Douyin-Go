package controller

import (
	"Douyin/common"
	"Douyin/dao"
	"Douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type FollowResponse struct {
	common.Response
	UserLists []dao.User `json:"user_list"`
}

func FollowAction(c *gin.Context) {
	//解析字段
	action := c.Query("action_type")
	toUserid, err := strconv.Atoi(c.Query("to_user_id"))
	userid := c.MustGet("userid").(int64)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	err = service.FollowOrCancel(userid, int64(toUserid), action)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  "successful",
	})
	return
}

func FollowList(c *gin.Context) {

	var userLists []dao.User
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FollowResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
		return
	}

	userLists, err = service.GetFollowLists(int64(userid))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FollowResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, FollowResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "successful"},
		UserLists: userLists,
	})
	return
}

func FollowerList(c *gin.Context) {

	var userLists []dao.User
	//预分配
	userLists = make([]dao.User, 0, 10)

	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FollowResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
		return
	}

	userLists, err = service.GetFollowerLists(int64(userid))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FollowResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, FollowResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "successful"},
		UserLists: userLists,
	})
	return
}
