package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Person 人
type Person struct {
	Name      string
	Phone     string
	City      string
	Age       int8
	IsMan     bool
	Interests []string
}

// One 数据结构
type One struct{
	ID bson.ObjectId `bson:"_id"` 
	Image string 
	Leyenda string
	Cita string
	Pubdate string
	Titulo string
}


//CreateData 创建数据
func CreateData(session *mgo.Session, dbname string, tablename string, data ...interface{}) error {

	cloneSession := session.Clone()
	c := cloneSession.DB(dbname).C(tablename)

	for _, item := range data {
		err := c.Insert(&item)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

// QueryAll 查询多条
func QueryAll(session *mgo.Session, dbname string, tablename string, query interface{}) *mgo.Iter {
    copySession := session.Clone()
    defer copySession.Close()

    collection := copySession.DB(dbname).C(tablename)

    //Using iterator prevent from taking up too much memory
    iter := collection.Find(query).Iter()

    if iter != nil {
        return iter
    }

    return nil
}