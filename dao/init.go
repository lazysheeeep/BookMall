package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var Db *gorm.DB

func DataBase(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "Debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{ //命名策略
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100) //打开连接数
	sqlDB.SetMaxIdleConns(20)  //最大连接池数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	Db = db

	//主从配置，读写分离
	_ = Db.Use(dbresolver.Register(
		dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connRead)},
			Replicas: []gorm.Dialector{mysql.Open(connWrite)},
			Policy:   dbresolver.RandomPolicy{},
		}))
	migration()
}

func NewDbClient(ctx context.Context) *gorm.DB {
	db := Db
	return db.WithContext(ctx)
}
