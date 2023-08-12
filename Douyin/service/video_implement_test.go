package service

import (
	"Douyin/cache"
	"Douyin/config"
	"Douyin/dao"
	"fmt"
	"testing"
)

func TestVideoStream(t *testing.T) {
	var err error
	err = config.ConfInit() //初始化配置文件
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = dao.DbInit() //初始化数据库
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cache.RedisPoolInit()
	t.Run("test1", func(t *testing.T) {
		token := "dasdadsadjnaidnd"
		videoInfo, _, err := VideoStream(token)
		if err != nil {
			t.Errorf("%v\n", err)
			return
		}
		fmt.Println(videoInfo)
	})
}
