package kits

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

type Mysql struct {
	DbName      string
	Host        string
	Port        string
	UserName    string
	Password    string
	MaxOpenConn int
	//指定了数据库连接池允许的最大同时打开的连接数
	MaxIdleConn int        //指定了数据库连接池中可以保持的最大空闲连接数。当前空闲连接数超过 MaxIdleConn 的限制，那么多余的连接将会被关闭
	MaxLifetime int64      //sec 指定了一个连接在被关闭之前的最大生命周期时间。如果一个连接的生命周期超过了 MaxLifetime 的限制，连接可能会被数据库服务器主动关闭。
	MaxIdleTime int64      //指定了一个连接在空闲状态下保持的最长时间
	l           sync.Mutex //锁
	DSN         string
}

var mysqlDb *gorm.DB

func NewMysql() *Mysql {
	m := &Mysql{
		MaxOpenConn: 100,
		MaxIdleConn: 10,
		MaxLifetime: 600,
		l:           sync.Mutex{},
		Host:        "mysql-container",
		Port:        "3306",
		UserName:    "lyc",
		Password:    "lyc",
		DbName:      "dora",
	}
	m.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
		m.UserName, m.Password, m.Host, m.Port, m.DbName)
	return m

}

func (m *Mysql) GetDB() (*gorm.DB, error) {
	m.l.Lock()
	defer m.l.Unlock()
	//使用锁来避免重复初始化,避免浪费资源

	if mysqlDb == nil {
		tmp, err := m.GetDBConn() //初始化一个连接池子，大家共用这个数据库对象和底层连接池
		if err != nil {
			return nil, err
		}
		mysqlDb = tmp
	}
	return mysqlDb, nil
}

func (m *Mysql) GetDBConn() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(m.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("mysql connect err: %v", err)
	}
	//	gorm.Open后,不会立即连接数据库，实际连接会在执行第一个查询时建立，ping会测试链接池中的链接，
	//连接不需要手动管理，连接池会自动管理连接的复用和释放
	sqldb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("mysql connect err: %v", err)
	}
	sqldb.SetMaxOpenConns(m.MaxOpenConn)
	sqldb.SetMaxIdleConns(m.MaxIdleConn)
	sqldb.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifetime))
	sqldb.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	//
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqldb.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql %s，error：%s", m.Host+"_"+m.Port, err.Error())
	}
	return db, nil
}
