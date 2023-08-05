package service

import (
	"Douyin/cache"
	"Douyin/common"
	"Douyin/dao"
	"Douyin/utils"
	"errors"
	"fmt"
)

func VideoStream(token string) ([]dao.Video, int64, error) {
	VideoLists, err := dao.GetVideoInstance().QueryVideo()
	//如果token不为空
	if token != "" {
		_, clime, _ := utils.ParseToken(token)
		var userLists []dao.User
		var likeVideoLists []dao.Video
		userLists, err = dao.GetFollowInstance().QueryFollowLists(clime.UserId)
		if err != nil {
			return VideoLists, 0, err
		}
		likeVideoLists, err = dao.GetLikeInstance().QueryLikeByUserid(clime.UserId)
		if err != nil {
			return VideoLists, 0, err
		}
		//fixme 缓存
		//设置 用户是否关注 视频是否点赞
		for index := range VideoLists {
			if cache.IsUserRelation(clime.UserId, VideoLists[index].Author.ID) {
				VideoLists[index].Author.IsFollow = true
			} else {
				for _, user := range userLists {
					if VideoLists[index].Author.ID == user.ID {
						VideoLists[index].Author.IsFollow = true
					}
				}
			}

			if cache.IsUserVideoRelation(clime.UserId, VideoLists[index].ID) {
				VideoLists[index].IsFavorite = true
			} else {
				for _, likeVideo := range likeVideoLists {
					if likeVideo.ID == VideoLists[index].ID {
						VideoLists[index].IsFavorite = true
					}
				}
			}

			err := common.UserCountSearchStrategy(&VideoLists[index].Author, VideoLists[index].Author.ID)
			if err != nil {
				return nil, 0, err
			}
		}
	}

	if len(VideoLists) == 0 {
		err = errors.New("video lists not exists")
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	//得到最早的时间返回过去
	nextTime := VideoLists[len(VideoLists)-1].TimeStamp
	return VideoLists, nextTime, nil

}
