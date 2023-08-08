package dao

import "sync"

type Comment struct {
	ID          int64  `gorm:"column:comment_id" json:"id,omitempty"`
	User        User   `gorm:"foreignKey:UserId"     json:"user"`
	UserId      int64  `gorm:"column:user_id"    json:"-"`
	VideoId     int64  `gorm:"column:video_id"   json:"-"`
	CommentText string `gorm:"column:comment_text"   json:"content,omitempty"`
	CreateTime  string `gorm:"column:create_time"    json:"create_date,omitempty"`
	TimeStamp   int64  `gorm:"column:timestamp"      json:"-"` //实现根据时间倒叙排列评论的功能
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

var commentDao *CommentDao

var commentOnce sync.Once

func GetCommentInstance() *CommentDao {
	commentOnce.Do(func() {
		commentDao = &CommentDao{}
	})
	return commentDao
}

//fixme:加入事务处理

func (CommentDao) AddComment(comment *Comment) error {
	tx := db.Begin() //开启事务
	err := tx.Create(comment).Error
	if err != nil {
		tx.Rollback() //错误则事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}
func (CommentDao) DeleteCommentById(commentId int64) error {
	tx := db.Begin() //开启事务
	//根据主键删除评论
	err := tx.Delete(&Comment{}, commentId).Error
	if err != nil {
		tx.Rollback() //错误则事务回滚
		return err
	}
	tx.Commit() //事务提交
	return nil
}

func (CommentDao) QueryListByVideoId(videoId int64) ([]Comment, error) {
	var commentLists []Comment
	tx := db.Begin() //开启事务
	//按时间的倒叙排序
	err := tx.Preload("User").Where("video_id =?", videoId).Order("timestamp desc").Find(&commentLists).Error
	if err != nil {
		tx.Rollback() //错误则事务回滚
		return nil, err
	}
	tx.Commit() //事务提交
	return commentLists, nil
}

func (CommentDao) QueryCommentByVideoId(videoId int64) ([]Comment, error) {
	var commentLists []Comment
	tx := db.Begin() //开启事务
	//按时间的倒叙排序
	err := tx.Preload("User").Where("video_id =?", videoId).Order("timestamp desc").Find(&commentLists).Error
	if err != nil {
		tx.Rollback() //错误则事务回滚
		return nil, err
	}
	tx.Commit() //事务提交
	return commentLists, nil
}
