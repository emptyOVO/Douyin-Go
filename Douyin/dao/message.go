package dao

import "sync"

type Message struct {
	ID         int64  `gorm:"column:msg_id"       json:"id,omitempty"`
	UserId     int64  `gorm:"column:from_user_id" json:"from_user_id,omitempty"`
	ToUserId   int64  `gorm:"column:to_user_id"   json:"to_user_id,omitempty"`
	Content    string `gorm:"column:content"      json:"content,omitempty"`
	CreateTime int64  `gorm:"column:create_time"  json:"create_time,omitempty"`
}

func (Message) TableName() string {
	return "message"
}

type MessageDao struct {
}

var messageDao *MessageDao
var messageOnce sync.Once

func GetMessageInstance() *MessageDao {
	messageOnce.Do(func() {
		messageDao = &MessageDao{}
	})
	return messageDao
}

// AddMessage 添加消息记录
func (MessageDao) AddMessage(message *Message) error {
	tx := db.Begin() //开启事务
	err := tx.Create(message).Error
	if err != nil {
		tx.Rollback() //事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

func (MessageDao) QueryMessageLists(userid, ToUserId int64, preMsgTime int64) ([]Message, error) {
	var MessageLists []Message
	err := db.Table("message").Where("create_time > ? AND ((to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?))", preMsgTime, userid, ToUserId, ToUserId, userid).Order("create_time").Find(&MessageLists).Error
	if err != nil {
		return nil, err
	}
	return MessageLists, nil
}
