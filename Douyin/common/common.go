package common

import (
	"Douyin/cache"
	"Douyin/dao"
)

type Response struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// UserCountSearchStrategy 计数的查找(获赞数，作品数，点赞数)
func UserCountSearchStrategy(user *dao.User, userid int64) error {
	var err error
	user.TotalFavorite, err = cache.GetUserTotalFavoriteCount(userid)
	if err != nil {
		user.TotalFavorite, err = dao.GetUserInstance().QueryTotalFavorite(userid)
		if err != nil {
			return err
		}
	}
	user.WorkCount, err = cache.GetUserWorkCount(userid)
	if err != nil {
		user.WorkCount, err = dao.GetUserInstance().QueryWorkCount(userid)
		if err != nil {
			return err
		}
	}
	user.FavoriteCount, err = cache.GetUserFavoriteCount(userid)
	if err != nil {
		user.FavoriteCount, err = dao.GetUserInstance().QueryFavoriteCount(userid)
		if err != nil {
			return err
		}
	}
	user.FollowerCount, err = cache.GetUserFollowerCount(userid)
	if err != nil {
		user.FollowerCount, err = dao.GetUserInstance().QueryFollowerCount(userid)
		if err != nil {
			return err
		}
	}
	user.FollowCount, err = cache.GetUserFollowCount(userid)
	if err != nil {
		user.FollowCount, err = dao.GetUserInstance().QueryFollowCount(userid)
		if err != nil {
			return err
		}
	}
	return nil
}
