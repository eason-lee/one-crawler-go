package main

import (
	_ "one-crawler-go/crawler"
	_ "one-crawler-go/routers"

	"github.com/astaxie/beego"
)

func main() {
	// collector := crawler.GetCollector()
	// // Start scraping on url
	// for i := 14; i <= 35; i++ {
	// 	collector.Visit("http://wufazhuce.com/one/" + strconv.Itoa(i))
	// }
	// collector.Wait()
	beego.Run()
}
