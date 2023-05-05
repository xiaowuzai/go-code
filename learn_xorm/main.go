package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type User struct {
	ID          string    `xorm:"pk varchar(36) notnull default('') comment('ID')"`
	Telephone   string    `xorm:"varchar(11) notnull default('') comment('手机号码')"` // 可以修改
	DownloadURL string    `xorm:"varchar(11) notnull default('') comment('下载链接')"` // 可以修改
	Description string    `xorm:"varchar(60) notnull default('') comment('描述')"`
	Password    []byte    `xorm:"blob comment('密码')" json:"password"`         // default: Admin@123
	Status      int32     `xorm:"notnull default(0) comment('状态:1.正常、2.禁用')"` // 1 正常、2 禁用
	ExpireTime  time.Time `xorm:"datetime notnull comment('债券名称')"`
	Created     time.Time `xorm:"datetime created"`
	Updated     time.Time `xorm:"datetime updated"`
}

func NewDB() (*xorm.Engine, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "AppStore123", "192.168.143.218", 3306, "mxshop_user_srv")
	engine, err := xorm.NewEngine("mysql", source)
	if err != nil {
		log.Printf("NewDB NewEngine error: %s", err.Error())
		return nil, err
	}

	engine.SetMapper(names.GonicMapper{})
	// 表前面统一加上 t_ 前缀
	tbMapper := names.NewPrefixMapper(names.GonicMapper{}, "t_")
	engine.SetTableMapper(tbMapper)

	// 设置时间
	engine.DatabaseTZ = time.UTC
	engine.TZLocation = time.UTC

	engine.ShowSQL(true)

	err = engine.Ping()
	if err != nil {
		logger.Error("NewDB Engine Ping error: ", err.Error())
		return nil, err
	}

	err = engine.Sync2(
		new(User),
	)
	if err != nil {
		logger.Error("NewDB Sync table error: ", err.Error())
		return nil, err
	}

	logger.Infof("NewDB success")
	return engine, nil
}

func main() {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}

	u := &User{
		ID:         "2",
		Telephone:  "2",
		ExpireTime: time.Unix(0, 0),
	}
	_, err = db.Insert(u)
	if err != nil {
		panic(err)
	}

	u1 := &User{}
	db.ID("2").Get(u1)
	fmt.Printf("%+v\n", u1)

	fmt.Println(u1.ExpireTime.Unix())
}
