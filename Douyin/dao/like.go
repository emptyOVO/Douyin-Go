package dao

import "sync"

type Like struct {
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

func (Like) TableName() string {
	return "like"
}

type LikeDao struct {
}

var likeDao *LikeDao

var likeOnce sync.Once

func GetLikeInstance() *LikeDao {
	likeOnce.Do(func() {
		likeDao = &LikeDao{}
	})
	return likeDao
}

// AddLike 增加
func (LikeDao) AddLike(like *Like) error {
	tx := db.Begin() //开启事务
	err := tx.Create(like).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

// DeleteLike 删除
func (LikeDao) DeleteLike(like *Like) error {
	tx := db.Begin() //开启事务
	err := tx.Where("user_id = ? and video_id = ?", like.UserId, like.VideoId).Delete(like).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

// QueryLikeByUserid DeleteLike 查找且返回lists
func (LikeDao) QueryLikeByUserid(userid int64) ([]Video, error) {
	user := &User{}
	tx := db.Begin() //开启事务
	err := tx.Preload("VideoLieLists").Where("user_id = ?", userid).Preload("VideoLieLists.Author").Find(user).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return nil, err
	}
	tx.Commit() //事务提交
	return user.VideoLieLists, nil
}
