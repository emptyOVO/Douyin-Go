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

type CommentResponse struct {
	common.Response
	Comment dao.Comment
}
type CommentListsResponse struct {
	common.Response
	CommentLists []dao.Comment `json:"comment_list"`
}

func CommentAction(c *gin.Context) {

	var (
		Comment   *dao.Comment
		videoId   int
		commentId int
		err       error
	)
	//解析各字段内容
	userid := c.MustGet("userid").(int64)
	action := c.Query("action_type")
	commentText := c.Query("comment_text")
	videoIdStr := c.Query("video_id")
	commentIdStr := c.Query("comment_id")

	if videoIdStr != "" {
		videoId, err = strconv.Atoi(c.Query("video_id"))
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, CommentResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
			})
			return
		}
	}
	if commentIdStr != "" {
		commentId, err = strconv.Atoi(commentIdStr)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, CommentResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
			})
			return
		}
	}
	//方法判断调用service
	Comment, err = service.CommentAddOrDelete(action, userid, int64(videoId), int64(commentId), commentText)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, CommentResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	if Comment == nil {
		c.JSON(http.StatusOK, CommentResponse{
			Response: common.Response{StatusCode: 0, StatusMsg: "successful"},
		})
		return
	}

	c.JSON(http.StatusOK, CommentResponse{
		Response: common.Response{StatusCode: 0, StatusMsg: "successful"},
		Comment:  *Comment,
	})
	return
}

func CommentList(c *gin.Context) {
	var commentLists []dao.Comment

	videoId, err := strconv.Atoi(c.Query("video_id"))

	if err != nil {
		c.JSON(http.StatusOK, CommentListsResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	commentLists, err = service.GetCommentLists(int64(videoId))
	if err != nil {
		c.JSON(http.StatusOK, CommentListsResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, CommentListsResponse{
		Response:     common.Response{StatusCode: 0, StatusMsg: "successful"},
		CommentLists: commentLists,
	})
	return
}
