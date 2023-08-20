package dao

import (
	"gorm.io/gorm"
	"log"
	"sync"
)

type User struct {
	ID              int64   `gorm:"column:user_id" json:"id"`
	Name            string  `gorm:"column:name"    json:"name"`
	FollowCount     int64   `gorm:"column:follow_count"   json:"follow_count"`
	FollowerCount   int64   `gorm:"column:follower_count" json:"follower_count"`
	Password        string  `gorm:"column:password"       json:"-"`
	IsFollow        bool    `gorm:"column:is_follow"                     json:"is_follow" `
	Avatar          string  `gorm:"column:avatar"         json:"avatar"`
	BackGroundImage string  `gorm:"column:background_image"        json:"background_image"`
	Signature       string  `gorm:"column:signature"               json:"signature"`
	TotalFavorite   int64   `gorm:"column:total_favorited"               json:"total_favorited"`
	WorkCount       int64   `gorm:"work_count"               json:"work_count"`
	FavoriteCount   int64   `gorm:"column:favorite_count"               json:"favorite_count"`
	VideoLieLists   []Video `gorm:"many2many:like;" json:"-"`
}

func (User) TableName() string { //表名
	return "user"
}

type UserDao struct {
}

var userDao *UserDao

// 只创建一次对象的单例接口
var userOnce sync.Once

func GetUserInstance() *UserDao {
	//创建单例
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

// AddUser 增加user
func (UserDao) AddUser(user *User) error {
	tx := db.Begin() //开启事务
	res := tx.Create(user)
	err := res.Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

// QueryUserByID user_id查找用户
func (UserDao) QueryUserByID(userID int64) (*User, error) {
	var user User
	tx := db.Begin() //开启事务
	err := tx.Where("user_id = ?", userID).Find(&user).Error
	if err != nil {
		tx.Rollback() //事务回滚
		log.Println(err.Error())
		return nil, err
	}
	tx.Commit() //事务提交
	return &user, nil
}

// QueryUserByName 名字查找用户
func (UserDao) QueryUserByName(name string) (*User, error) {
	var user User
	tx := db.Begin() //开启事务
	err := tx.Where("name = ?", name).Find(&user).Error
	if err != nil {
		tx.Rollback() //事务回滚
		log.Println(err.Error())
		return nil, err
	}
	tx.Commit() //事务提交
	return &user, nil
}

// UpdateFollowCount 更新关注数量
func (UserDao) UpdateFollowCount(userId, count int64) error {
	tx := db.Begin() //开启事务
	err := tx.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", count)).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

// UpdateTotalFavoriteCount 更新获赞数量
func (UserDao) UpdateTotalFavoriteCount(userId, count int64) error {
	tx := db.Begin() //开启事务
	err := tx.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("total_favorited", gorm.Expr("total_favorited + ?", count)).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

// UpdateFavoriteCount 更新点赞数量
func (UserDao) UpdateFavoriteCount(userId, count int64) error {
	tx := db.Begin() //开启事务
	err := tx.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", count)).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

// UpdateFollowerCount 更新粉丝数量
func (UserDao) UpdateFollowerCount(userId, count int64) error {
	tx := db.Begin() //开启事务
	err := tx.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", count)).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

// UpdateWorkCount 更新作品数量
func (UserDao) UpdateWorkCount(userId, count int64) error {
	tx := db.Begin() //开启事务
	err := tx.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("work_count", gorm.Expr("work_count + ?", count)).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

// QueryWorkCount 获取作品数量
func (UserDao) QueryWorkCount(userid int64) (int64, error) {
	var count int64
	tx := db.Begin() //开启事务
	err := tx.Model(&Video{}).Where("user_id = ?", userid).Count(&count).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return 0, err
	}
	tx.Commit() //事务提交
	return count, nil
}

// QueryFavoriteCount 获取获赞的总数量
func (UserDao) QueryFavoriteCount(userid int64) (int64, error) {
	var count int64
	tx := db.Begin()
	defer tx.Commit()
	err := tx.Model(&Like{}).Where("user_id = ?", userid).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return count, nil
}

// QueryTotalFavorite 获取收获的赞总数
func (UserDao) QueryTotalFavorite(userid int64) (int64, error) {
	var count int64
	tx := db.Begin()
	defer tx.Commit()
	err := tx.Model(&Like{}).
		Select("COUNT(*)").
		Joins("JOIN video ON like.video_id = video.video_id").
		Where("video.user_id = ?", userid).
		Count(&count).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return count, nil
}

// QueryFollowCount 获取关注的用户数
func (UserDao) QueryFollowCount(userid int64) (int64, error) {
	var count int64
	tx := db.Begin()
	defer tx.Commit()
	err := tx.Model(&Follow{}).Where("follow_id = ?", userid).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return count, nil
}

// QueryFollowerCount 获取粉丝数
func (UserDao) QueryFollowerCount(userid int64) (int64, error) {
	var count int64
	tx := db.Begin()
	defer tx.Commit()
	err := tx.Model(&Follow{}).Where("followed_id = ?", userid).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return count, nil
}
