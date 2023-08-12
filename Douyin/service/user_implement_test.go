package service

import (
	"Douyin/cache"
	"Douyin/config"
	"Douyin/dao"
	"fmt"
	"testing"
)

func TestUserRegister(t *testing.T) {
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
	err = cache.RedisPoolInit()
	if err != nil {
		t.Error(err.Error())
		return
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args
	}{
		{
			"测试小王",
			args{
				"____________",
				"123456mksjxnjancanjskandndjasnjdkasn",
			},
		},
		{
			"测试小贺",
			args{
				"小王",
				"123456",
			},
		},
		{
			"测试小陈",
			args{
				"",
				"123456",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userinfo, err := UserRegister(test.username, test.password)
			if err != nil {
				t.Errorf("UserRegister ERROR is %v", err)
				return
			}
			fmt.Println(userinfo)
		})
	}
}

func TestUserLogin(t *testing.T) {
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
		username string
		password string
	}
	tests := []struct {
		name string
		args
	}{
		{
			"测试1",
			args{
				"小王",
				"123456",
			},
		},
		{
			"测试2",
			args{
				"_____",
				"123456",
			},
		},
		{
			"测试3",
			args{
				"hhhy",
				"123456",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userinfo, err := UserLogin(test.username, test.password)
			if err != nil {
				t.Errorf("UserRegister ERROR is %v", err)
				return
			}
			fmt.Println(userinfo)
		})
	}
}
