package dao

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type Video struct {
	ID            int64  `gorm:"column:video_id"       json:"id,omitempty"`
	Author        User   `gorm:"foreignKey:UserId"     json:"author"`
	UserId        int64  `gorm:"column:user_id"        json:"-"`
	PlayUrl       string `gorm:"column:play_url"       json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:cover_url"      json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count"  json:"comment_count,omitempty"`
	Title         string `gorm:"column:title"          json:"title,omitempty"`
	TimeStamp     int64  `gorm:"column:timestamp"      json:"-"`
	IsFavorite    bool   `gorm:"column:is_favorite"                     json:"is_favorite"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func GetVideoInstance() *VideoDao {
	//创建单例
	videoOnce.Do(func() {
		videoDao = &VideoDao{}
	})
	return videoDao
}

// QueryVideo 初始化 视频流
func (VideoDao) QueryVideo() ([]Video, error) {
	var videoLists []Video
	//预分配
	videoLists = make([]Video, 0, 10)
	tx := db.Begin() //开启事务
	//查询所有的视频 及其作者信息 及按时间的降序进行排列
	err := tx.Preload("Author").Order("timestamp desc").Find(&videoLists).Error
	if err != nil {
		tx.Rollback() //事务回滚
		fmt.Println(err.Error())
		return nil, err
	}
	tx.Commit() //事务提交
	return videoLists, nil
}

func (VideoDao) AddVideo(video *Video) error {
	tx := db.Begin() //开启事务
	res := tx.Create(video)
	err := res.Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

func (VideoDao) UpdateCommentCount(videoId, count int64) error {
	tx := db.Begin() //开启事务
	err := tx.Model(&Video{}).Where("video_id = ?", videoId).UpdateColumn("comment_count", gorm.Expr(" comment_count + ?", count)).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}
func (VideoDao) UpdateFavoriteCount(videoId, count int64) error {
	tx := db.Begin() //开启事务
	err := tx.Model(&Video{}).Where("video_id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr(" favorite_count + ?", count)).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

func (VideoDao) QueryUserIdByVideoId(videoId int64) (int64, error) {
	var userid int64
	tx := db.Begin() //开启事务
	err := tx.Raw("SELECT user_id FROM video WHERE video.video_id = ?", videoId).Scan(&userid).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return 0, err
	}
	tx.Commit() //事务提交
	return userid, nil
}

func (VideoDao) QueryVideoByUserId(userid int64) ([]Video, error) {
	var videoLists []Video
	//预分配
	videoLists = make([]Video, 0, 10)
	tx := db.Begin() //开启事务
	//查询用户发布的视频
	err := tx.Preload("Author").Where("video.user_id = ?", userid).Order("timestamp desc").Find(&videoLists).Error
	if err != nil {
		tx.Rollback() //事务回滚
		fmt.Println(err.Error())
		return nil, err
	}
	tx.Commit() //事务提交
	return videoLists, nil
}
