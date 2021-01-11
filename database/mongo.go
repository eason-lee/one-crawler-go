package db

import (
	"log"
	"one-crawler-go/config"
	"time"

	"gopkg.in/mgo.v2"
)

// GlobalDatabase 全局 session
var GlobalDatabase *mgo.Session

// Init 初始化 mongo
func Init(conf *config.MongoConf) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{conf.Hosts},
		Timeout:  60 * time.Second,
		Database: conf.DB,
		// Username: AuthUserName,
		// Password: AuthPassword,
	}

	var err error
	GlobalDatabase, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession failed:%\n", err)
	}

	//设置连接池的大小
	GlobalDatabase.SetPoolLimit(conf.MaxCon)
}
