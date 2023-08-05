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
	res := db.Create(user)
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}

// QueryUserByID 通过id查找user
func (UserDao) QueryUserByID(userID int64) (*User, error) {
	var user User
	err := db.Where("user_id = ?", userID).Find(&user).Error
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// QueryUserByName 名字查找user
func (UserDao) QueryUserByName(name string) (*User, error) {
	var user User
	err := db.Where("name = ?", name).Find(&user).Error
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateFollowCount 更新关注数量
func (UserDao) UpdateFollowCount(userId, count int64) error {
	err := db.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", count)).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateTotalFavoriteCount 更新获赞数量
func (UserDao) UpdateTotalFavoriteCount(userId, count int64) error {
	err := db.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("total_favorited", gorm.Expr("total_favorited + ?", count)).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateFavoriteCount 更新点赞数量
func (UserDao) UpdateFavoriteCount(userId, count int64) error {
	err := db.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", count)).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateFollowerCount 更新粉丝数量
func (UserDao) UpdateFollowerCount(userId, count int64) error {
	err := db.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", count)).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateWorkCount 更新作品数量
func (UserDao) UpdateWorkCount(userId, count int64) error {
	err := db.Model(&User{}).Where("user_id = ?", userId).UpdateColumn("work_count", gorm.Expr("work_count + ?", count)).Error
	if err != nil {
		return err
	}
	return nil
}

// QueryWorkCount 获取作品数量
func (UserDao) QueryWorkCount(userid int64) (int64, error) {
	var count int64
	err := db.Raw("SELECT COUNT(*) FROM video WHERE video.user_id = ?", userid).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// QueryFavoriteCount 获赞的总数量
func (UserDao) QueryFavoriteCount(userid int64) (int64, error) {
	var count int64
	err := db.Raw("SELECT COUNT(*) FROM `like` WHERE `like`.user_id = ?", userid).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// QueryTotalFavorite  获取收获的赞总数
func (UserDao) QueryTotalFavorite(userid int64) (int64, error) {
	var count int64
	err := db.Raw("SELECT COUNT(*) FROM `like` WHERE `like`.video_id in (SELECT video_id FROM video WHERE video.user_id  = ? )", userid).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// QueryFollowCount  获取关注的用户数
func (UserDao) QueryFollowCount(userid int64) (int64, error) {
	var count int64
	err := db.Raw("SELECT COUNT(*) FROM `follow` WHERE `follow`.follow_id = ?", userid).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// QueryFollowerCount  获取粉丝数
func (UserDao) QueryFollowerCount(userid int64) (int64, error) {
	var count int64
	err := db.Raw("SELECT COUNT(*) FROM `follow` WHERE `follow`.followed_id = ?", userid).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
