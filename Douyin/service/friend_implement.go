package service

import (
	"Douyin/common"
	"Douyin/dao"
	"errors"
	"fmt"
)

func GetFriendLists(userid int64) ([]dao.User, error) {
	var err error
	var UserLists []dao.User
	UserLists, err = dao.GetFollowInstance().QueryEachFollow(userid)
	//将状态设置为已经关注
	for index := range UserLists {
		UserLists[index].IsFollow = true
		err := common.UserCountSearchStrategy(&UserLists[index], UserLists[index].ID)
		if err != nil {
			return nil, err
		}
	}
	if len(UserLists) == 0 {
		err = errors.New("users not exists")
		return nil, err
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return UserLists, nil
}
