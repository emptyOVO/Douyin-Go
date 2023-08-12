package service

import (
	"Douyin/cache"
	"Douyin/config"
	"Douyin/dao"
	"fmt"
	"testing"
)

func TestPublishedVideoLists(t *testing.T) {
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
	type args struct {
		userid int64
	}
	tests := []struct {
		name string
		args
	}{
		{
			"测试1",
			args{
				1,
			},
		},
		{
			"测试2",
			args{
				100000,
			},
		},
		{
			"测试3",
			args{
				-10,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			videoInfo, err := PublishedVideoLists(test.userid)
			if err != nil {
				t.Errorf("%v\n", err)
				return
			}
			fmt.Printf("%#v", videoInfo)
		})
	}
}

func TestPublishVideo(t *testing.T) {
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
	type args struct {
		userid   int64
		playUrl  string
		coverUrl string
		title    string
	}
	tests := []struct {
		name string
		args
	}{
		{
			name: "测试1",
			args: args{
				userid:   1,
				playUrl:  "http://192.168.2.166:80/public/1_wx_camera_1690344007531.mp4",
				coverUrl: "http://192.168.2.166:80/public/1_wx_camera_1690344007531.mp4.png",
				title:    "1",
			},
		},
		{
			name: "测试2",
			args: args{
				userid:   2,
				playUrl:  "http://192.168.2.166:80/public/2_mmexport1690506409070.mp4",
				coverUrl: "http://192.168.2.166:80/public/2_mmexport1690506409070.mp4.png",
				title:    "3",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := PublishVideo(test.userid, test.playUrl, test.coverUrl, test.title)
			fmt.Printf("%#v", err)
			if err != nil {
				t.Errorf("ERROR is %v", err)
				return
			}
		})
	}

}
