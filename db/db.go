/**
 * Created by lock
 * Date: 2019-09-22
 * Time: 22:37
 */
package db

import (
	"gochat/config"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbMap = map[string]*gorm.DB{}
var syncLock sync.Mutex

func init() {
	initDB("gochat")
}

func getLoggerByConfig() logger.Interface {
	// 从配置文件获取日志级别
	logLevel := "warn"
	if config.GetMode() == "dev" {
		logLevel = "info"
	}
	switch logLevel {
	case "error":
		return logger.Default.LogMode(logger.Error)
	case "warn":
		return logger.Default.LogMode(logger.Warn)
	case "info":
		return logger.Default.LogMode(logger.Info)
	default:
		return logger.Default.LogMode(logger.Info)
	}
}
func initDB(dbName string) {
	var e error
	// if prod env , you should change mysql driver for yourself !!!
	// realPath, _ := filepath.Abs("./")
	// configFilePath := realPath + "/db/gochat.sqlite3"
	syncLock.Lock()
	// dbMap[dbName], e = gorm.Open("sqlite3", configFilePath)
	dsn := "root:weile2017i@tcp(192.168.67.9:3306)/gozero"

	dbMap[dbName], e = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: getLoggerByConfig(),
	})

	if e != nil {
		panic(e)
	}
	sqlDB, _ := dbMap[dbName].DB()
	sqlDB.SetMaxIdleConns(4)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(8 * time.Second)

	sqlDB.SetMaxIdleConns(4)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(8 * time.Second)

	syncLock.Unlock()
	if e != nil {
		logrus.Error("connect db fail:%s", e.Error())
	}
}

func GetDb(dbName string) (db *gorm.DB) {
	if db, ok := dbMap[dbName]; ok {
		return db
	} else {
		return nil
	}
}

type DbGoChat struct {
}

func (*DbGoChat) GetDbName() string {
	return "gochat"
}
