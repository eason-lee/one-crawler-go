package config

// AppConf App配置
type AppConf struct {
	CrawlerConf   `yaml:"crawler"`
	MongoConf    `yaml:"mongo"`
}

// CrawlerConf ...
type CrawlerConf struct {
	// 是否开启异步
	Async     bool `yaml:"async"`
	// 爬虫的起点
	StartIndex int      `yaml:"start_index"`
	// 爬虫的终点
	EndIndex int      `yaml:"end_index"`
}

// MongoConf ...
type MongoConf struct {
	Hosts string `yaml:"hosts"`
	DB string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	MaxCon int `yaml:"max_connect"`
}

