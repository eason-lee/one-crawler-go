package controllers

import (
	db "one-crawler-go/database"
	"one-crawler-go/models"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

// IndexController 启动页控制器
type IndexController struct {
	beego.Controller
}

const pageSize = 5

// Index 页面启动
func (c *IndexController) Index() {
	iter := db.GlobalDatabase.DB("test").C("one").Find(bson.M{}).Sort("-_id").Limit(pageSize).Iter()
	result := models.One{}
	results := []models.One{}
	for iter.Next(&result) {
		results = append(results, result)
	}
	c.Data["FirstOne"] = results[len(results)-1].ID.Hex()
	c.Data["LastOne"] = results[0].ID.Hex()
	c.Data["Ones"] = results
	c.TplName = "ones.tpl"
}

// PreOneView one 上一页
func (c *IndexController) PreOneView() {
	id := c.Ctx.Input.Param(":id")

	iter := db.GlobalDatabase.DB("test").C("one").Find(
		bson.M{"_id": bson.M{"$lt": bson.ObjectIdHex(id)}}).Limit(pageSize).Sort("-_id").Iter()
	result := models.One{}
	results := []models.One{}
	for iter.Next(&result) {
		results = append(results, result)
	}
	if len(results) == 0 {
		iter = db.GlobalDatabase.DB("test").C("one").Find(
			bson.M{}).Sort("_id").Limit(pageSize).Sort("-_id").Iter()
		for iter.Next(&result) {
			results = append(results, result)
		}
	}

	c.Data["Ones"] = results
	c.Data["FirstOne"] = results[len(results)-1].ID.Hex()
	c.Data["LastOne"] = results[0].ID.Hex()
	c.TplName = "ones.tpl"
}

// NextOneView one 下一页
func (c *IndexController) NextOneView() {
	id := c.Ctx.Input.Param(":id")

	iter := db.GlobalDatabase.DB("test").C("one").Find(
		bson.M{"_id": bson.M{"$gt": bson.ObjectIdHex(id)}},
	).Sort("_id").Limit(pageSize).Sort("-_id").Iter()

	result := models.One{}
	results := []models.One{}
	for iter.Next(&result) {
		results = append(results, result)
	}
	if len(results) == 0 {
		iter = db.GlobalDatabase.DB("test").C("one").Find(
			bson.M{}).Sort("-_id").Limit(pageSize).Iter()
		for iter.Next(&result) {
			results = append(results, result)
		}
	}
	c.Data["FirstOne"] = results[len(results)-1].ID.Hex()
	c.Data["LastOne"] = results[0].ID.Hex()
	c.Data["Ones"] = results
	c.TplName = "ones.tpl"
}
