package crawler

import (
	"fmt"
	db "one-crawler-go/database"
	"one-crawler-go/config"
	"one-crawler-go/models"
	"strconv"

	"github.com/gocolly/colly/v2"
	"gopkg.in/mgo.v2/bson"
)

// getCollector 获取 Collector
func getCollector(async bool) *colly.Collector {
	c := colly.NewCollector(

	// 开启异步
	colly.Async(async),
	)

	// On every a element which has href attribute call callback
	// one 图片
	c.OnHTML("body", func(e *colly.HTMLElement) {
		one := new(models.One)

		image := e.ChildAttr(".one-imagen>img", "src")

		// fmt.Printf("one image found: %q ->\n",  image)
		one.Image = image

		cita := e.ChildText(".one-cita")
		// fmt.Printf("one intro found: %q -> \n", cita)
		one.Cita = cita

		leyenda := e.ChildText(".one-imagen-leyenda")
		// fmt.Printf("one leyenda found: %q -> \n", leyenda)
		one.Leyenda = leyenda

		titulo := e.ChildText(".one-titulo")
		// fmt.Printf("one titulo found: %q -> \n", titulo)
		one.Titulo = titulo

		pubdateDay := e.ChildText(".one-pubdate>.dom")
		pubdateMonth := e.ChildText(".one-pubdate>.may")
		pubdate := pubdateDay + " " + pubdateMonth
		// fmt.Printf("one titulo found: %q -> \n", pubdate)
		one.Pubdate = pubdate
		one.ID = bson.NewObjectId()

		data := []models.One{}
		data = append(data, *one)

		fmt.Printf("one data %#v \n", data)
		err := models.CreateData(db.GlobalDatabase, "test", "one", one)
		if err != nil {
			panic(err)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	return c

}

// Run 开启爬虫
func Run(conf *config.CrawlerConf) {
	collector := getCollector(conf.Async)
	// Start scraping on url
	for i := conf.StartIndex; i <= conf.EndIndex; i++ {
		collector.Visit("http://wufazhuce.com/one/" + strconv.Itoa(i))
	}
	if conf.Async{
		collector.Wait()
	}
}
