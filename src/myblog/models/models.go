package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

// 用户
type User struct {
	Id       int
	Username string
	Password string
}

// 文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

const (
	_MYSQL_DRIVER = "mysql"
)

func RegisterDB() {
	orm.RegisterModel(new(User), new(Topic))
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DR_MySQL)
	orm.RegisterDataBase("default", _MYSQL_DRIVER, "root:feng@/blog?charset=utf8", 30)
}

// 添加文章
func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}

// 获取所有文章
func GetAllTopic(isDesc bool, p int) (topic []*Topic, err error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	pageNum, _ := beego.AppConfig.Int("pageNum")
	if isDesc {
		_, err = qs.OrderBy("-Created").Limit(pageNum, (p-1)*pageNum).All(&topics)
	} else {
		_, err = qs.Limit(pageNum, (p-1)*pageNum).All(&topics)
	}
	return topics, err
}

// 获取单个文章信息
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

// 获取文章总数
func GetTopicCount(title string) (count int64, err error) {
	o := orm.NewOrm()
	if len(title) > 0 {
		count, err = o.QueryTable("topic").Filter("title__icontains", title).Count()
	} else {
		count, err = o.QueryTable("topic").Count()
	}

	if err != nil {
		return 0, err
	}
	fmt.Println("count: ", count)
	return count, err
}

// 文章搜索
func SearchTopic(title string) (topic []*Topic, err error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	_, err = o.QueryTable("topic").Filter("title__icontains", title).All(&topics)
	return topics, err
}
