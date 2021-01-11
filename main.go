package main

import (
	"io/ioutil"
	"log"
	"one-crawler-go/config"
	db "one-crawler-go/database"
	"one-crawler-go/crawler"
	_ "one-crawler-go/routers"

	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
)

func main() {
	// 从配置文件初始化 server 配置项
	conf := new(config.AppConf)
	yamlFile, err := ioutil.ReadFile("config/app.yaml")
	if err != nil {
		log.Fatalf("yamlFile Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("yaml conf Unmarshal: %v", err)
	}

	// 初始化 mongo
	db.Init(&conf.MongoConf)
	// 开启爬虫
	go crawler.Run(&conf.CrawlerConf)
	// web run
	log.Println("开启 beego")
	beego.Run()
}
