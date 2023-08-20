package dao

import (
	"fmt"
	"sync"
)

type Follow struct {
	FollowId   int64 `gorm:"column:follow_id"`
	FollowedId int64 `gorm:"column:followed_id"`
}

func (Follow) TableName() string {
	return "follow"
}

type FollowDao struct {
}

var followDao *FollowDao

var FollowOnce sync.Once

func GetFollowInstance() *FollowDao {
	FollowOnce.Do(func() {
		followDao = &FollowDao{}
	})
	return followDao
}

func (FollowDao) QueryAllFollow() ([]Follow, error) {

	var FollowLists []Follow
	tx := db.Begin() //开启事务
	err := tx.Find(&FollowLists).Error
	if err != nil {
		tx.Rollback() //错误则事务回滚
		return nil, err
	}
	tx.Commit() //事务提交
	return FollowLists, err
}

// AddFollow 添加关注映射
func (FollowDao) AddFollow(follow *Follow) error {
	tx := db.Begin() //开启事务
	err := tx.Create(follow).Error
	if err != nil {
		tx.Rollback() //错误则事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

// DeleteFollow 删除关注映射
func (FollowDao) DeleteFollow(follow *Follow) error {
	tx := db.Begin() //开启事务
	err := tx.Where("follow_id = ? and followed_id = ?", follow.FollowId, follow.FollowedId).Delete(follow).Error
	if err != nil {
		tx.Rollback() //错误则事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

func (FollowDao) QueryFollowLists(userid int64) ([]User, error) {
	var userLists []User
	err := db.Table("user").Where("user_id IN (?)", db.Table("follow").Select("followed_id").Where("follow_id = ?", userid)).Scan(&userLists).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v", userLists)
	return userLists, nil
}
func (FollowDao) QueryFollowerLists(userid int64) ([]User, error) {
	var userLists []User
	err := db.Table("user").Where("user_id IN (?)", db.Table("follow").Select("follow_id").Where("followed_id = ?", userid)).Scan(&userLists).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v", userLists)
	return userLists, nil
}
func (FollowDao) QueryEachFollow(userid int64) ([]User, error) {
	var userLists []User
	subQuery := db.Table("follow").Select("follow_id").Where("followed_id = ?", userid)
	err := db.Table("user").Where("user_id != ? AND user_id IN (?)", userid, db.Table("follow").Select("DISTINCT followed_id").Joins("JOIN (?) a ON a.follow_id = follow.followed_id", subQuery)).Scan(&userLists).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v", userLists)
	return userLists, nil
}
