package service

import (
	"Douyin/cache"
	"Douyin/config"
	"Douyin/dao"
	"fmt"
	"testing"
)

func TestGetFriendLists(t *testing.T) {
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
	tests := []struct {
		name   string
		userId int64
	}{
		{
			name:   "test1",
			userId: 10,
		},
	}
	for _, test := range tests {
		t.Run("测试", func(t *testing.T) {
			commentLists, err := GetFriendLists(test.userId)
			if err != nil {
				t.Errorf("%#v", err)
				return
			}
			fmt.Printf("%#v", commentLists)
		})

	}
}
