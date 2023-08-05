package service

import (
	"Douyin/cache"
	"Douyin/common"
	"Douyin/dao"
	"errors"
	"fmt"
	"log"
	"time"
)

// PublishedVideoLists 已发布视频的列表
func PublishedVideoLists(userid int64) ([]dao.Video, error) {
	var (
		VideoLists []dao.Video
		err        error
	)
	VideoLists, err = dao.GetVideoInstance().QueryVideoByUserId(userid)
	if len(VideoLists) == 0 {
		err = errors.New("video lists not exists")
		return nil, err
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//缓存层
	for index1 := range VideoLists {
		if cache.IsUserVideoRelation(userid, VideoLists[index1].ID) {
			VideoLists[index1].IsFavorite = true
		}
		err := common.UserCountSearchStrategy(&VideoLists[index1].Author, userid)
		if err != nil {
			return nil, err
		}
	}
	if len(VideoLists) >= 30 {
		//限制只能返回最多三十条视频
		return VideoLists[0:30], nil
	}

	return VideoLists, nil
}

// PublishVideo 投稿视频
func PublishVideo(userid int64, playUrl, coverUrl, title string) error {
	err := dao.GetVideoInstance().AddVideo(&dao.Video{
		UserId:        userid,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		TimeStamp:     time.Now().Unix(),
		IsFavorite:    false,
	})
	if err != nil {
		return err
	}
	err = cache.IncrByUserWorkCount(userid)
	if err != nil {
		log.Println(err.Error())
	}
	// 并且更新作品数量
	err = dao.GetUserInstance().UpdateWorkCount(userid, 1)
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}
