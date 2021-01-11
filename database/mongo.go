package db

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

const (
	mongoDBHosts = "127.0.0.1:27017"
	authDatabase = "test"
	authUserName = "test"
	authPassword = "123456"
	maxCon       = 300
)

// GlobalDatabase 全局 session
var GlobalDatabase = getGlobalSession()

func getGlobalSession() *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{mongoDBHosts},
		Timeout:  60 * time.Second,
		Database: authDatabase,
		// Username: AuthUserName,
		// Password: AuthPassword,
	}

	GlobalDatabase, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession failed:%\n", err)
	}

	//设置连接池的大小
	GlobalDatabase.SetPoolLimit(maxCon)

	return GlobalDatabase
}
