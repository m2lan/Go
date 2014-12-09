package models

import (
	"fmt"
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
	orm.RegisterDataBase("default", _MYSQL_DRIVER, "root:root@/blog?charset=utf8", 30)
}

func SelectList() {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("SELECT username FROM user WHERE id = ?", 1).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["username"]) // slene
	}
}

func SelectList2() {
	o := orm.NewOrm()
	u := User{Id: 1}
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		fmt.Println("没有数据")
	} else if err == orm.ErrMissPK {
		fmt.Println("没有主键")
	} else {
		fmt.Println(u.Username, u.Password)
	}
}

func SelectList3() {
	o := orm.NewOrm()
	u := Topic{Id: 1}
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		fmt.Println("没有数据")
	} else if err == orm.ErrMissPK {
		fmt.Println("没有主键")
	} else {
		fmt.Println(u.Title, u.Content)
	}
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
func GetAllTopic(isDesc bool) (topic []*Topic, err error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	if isDesc {
		_, err = qs.OrderBy("-Created").All(&topics)
	} else {
		_, err = qs.All(&topics)
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
