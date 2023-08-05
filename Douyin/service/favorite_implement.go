package service

import (
	"Douyin/cache"
	"Douyin/common"
	"Douyin/dao"
	"errors"
	"fmt"
	"log"
)

// GetLikeLists  返回喜欢列表
func GetLikeLists(userid int64) ([]dao.Video, error) {
	VideoLists, err := dao.GetLikeInstance().QueryLikeByUserid(userid)
	if len(VideoLists) == 0 {
		err = errors.New("video lists not exist")
		return nil, err
	}
	for index := range VideoLists {
		VideoLists[index].IsFavorite = true
		err := common.UserCountSearchStrategy(&VideoLists[index].Author, VideoLists[index].Author.ID)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return VideoLists, nil
}

func LikeOrCancel(FavoriteAction string, userid int64, videoId int64) error {
	var (
		err    error
		userId int64
	)

	userId, err = dao.GetVideoInstance().QueryUserIdByVideoId(videoId)
	if err != nil {
		fmt.Println(err.Error())
	}

	//点赞和取消点赞的动作 更新数量
	switch FavoriteAction {
	//加入redis缓存
	case "1":
		err = dao.GetLikeInstance().AddLike(&dao.Like{
			UserId:  userid,
			VideoId: videoId,
		})
		if err != nil {
			return err
		}
		err = dao.GetVideoInstance().UpdateFavoriteCount(videoId, 1)
		if err != nil {
			return err
		}
		err = dao.GetUserInstance().UpdateFavoriteCount(userid, 1)
		if err != nil {
			return err
		}
		err = cache.SetUserVideoRelation(userid, videoId)
		if err != nil {
			log.Println(err.Error())
		}
		//缓存加一
		err = cache.IncrByUserFavoriteCount(userid)
		if err != nil {
			//fixme 这里不应该直接返回错误 因为只是缓存失效了
			fmt.Println(err.Error())
		}
		//缓存加一
		err = cache.IncrByUserTotalFavorite(userId)
		//更新total_favorite计数
		err = dao.GetUserInstance().UpdateTotalFavoriteCount(userId, 1)
		if err != nil {
			//fixme 这里不应该直接返回err 仅仅只是缓存
			fmt.Println(err.Error())
		}
	case "2":
		err = dao.GetLikeInstance().DeleteLike(&dao.Like{
			UserId:  userid,
			VideoId: videoId,
		})
		if err != nil {
			return err
		}
		err = dao.GetVideoInstance().UpdateFavoriteCount(videoId, -1)
		if err != nil {
			return err
		}
		err = dao.GetUserInstance().UpdateFavoriteCount(userid, -1)
		if err != nil {
			return err
		}
		//缓存减一
		err = cache.DecrByUserFavoriteCount(userid)
		if err != nil {
			//fixme 这里不应该直接返回错误 因为只是缓存失效了
			log.Println(err.Error())
		}
		//缓存减一
		err = cache.DecrByUserTotalFavorite(userId)
		//更新total_favorite计数
		err = dao.GetUserInstance().UpdateTotalFavoriteCount(userId, -1)
		if err != nil {
			return err
		}
		err = cache.DeleteUserVideoRelation(userid, videoId)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return nil
}
