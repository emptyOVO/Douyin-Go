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

type MessageResponse struct {
	common.Response
	MessageLists []dao.Message `json:"message_list"`
}

// MessageAction 消息发送接口
func MessageAction(c *gin.Context) {
	//解析各字段内容
	userid := c.MustGet("userid").(int64)
	toUserId, err := strconv.Atoi(c.Query("to_user_id"))
	action := c.Query("action_type")
	content := c.Query("content")
	//错误判断
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//service方法调用
	err = service.SendMessage(userid, int64(toUserId), action, content)
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

// MessageChat 获取消息记录
func MessageChat(c *gin.Context) {
	var (
		toUserId     int
		preMsgTime   int
		err          error
		messageLists []dao.Message
	)
	//得到userid
	userid := c.MustGet("userid").(int64)
	//得到对方用户id
	toUserId, err = strconv.Atoi(c.Query("to_user_id"))
	//得到上次最新消息的时间
	preMsgTime, err = strconv.Atoi(c.Query("pre_msg_time"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, MessageResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, MessageResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	messageLists, err = service.GetMessageLists(userid, int64(preMsgTime), int64(toUserId))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, MessageResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "successful",
		},
		MessageLists: messageLists,
	})
	return
}
