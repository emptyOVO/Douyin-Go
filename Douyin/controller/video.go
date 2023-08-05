package controller

import (
	"Douyin/common"
	"Douyin/dao"
	"Douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FeedResponse struct {
	common.Response
	VideoLists []dao.Video `json:"video_list,omitempty"`
	NextTime   int64       `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	token := c.Query("token")
	//返回视频信息
	//按照时间降序
	videoLists, timestamp, err := service.VideoStream(token)
	if len(videoLists) >= 30 {
		//限制返回最多三十条结果
		videoLists = videoLists[0:30]
	}
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FeedResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response:   common.Response{StatusCode: 0, StatusMsg: "successful"},
			VideoLists: videoLists,
			NextTime:   timestamp,
		})
		return
	}
}
