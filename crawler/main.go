package crawler

import (
	"fmt"
	// "log"
	db "one-crawler-go/database"
	"one-crawler-go/models"

	"github.com/gocolly/colly/v2"
	"gopkg.in/mgo.v2/bson"
)

// GetCollector 获取 Collector
func GetCollector() *colly.Collector {
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// colly.AllowedDomains("http://wufazhuce.com"),
	// 开启异步
	// colly.Async(true),
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
		err :=  models.CreateData(db.GlobalDatabase, "test", "one", one)
		if err != nil{
			panic(err)
		}
	})


	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	return c

}
