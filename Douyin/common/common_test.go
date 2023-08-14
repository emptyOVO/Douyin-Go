package common

import (
	"Douyin/config"
	"Douyin/dao"
	"fmt"
	"testing"
)

func TestUserCountSearchStrategy(t *testing.T) {
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
	user, err := dao.GetUserInstance().QueryUserByID(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = UserCountSearchStrategy(user, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
