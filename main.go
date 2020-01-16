package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	p1 := Line{
		// ID:      1,
		Name:    "黄洁",
		Account: "@vfc2617i",
		Fans:    5134,
		Paas:    1,
		Status:  1,
	}
	op := new(Operater)
	op.dbname = "dsp"
	op.document = "line"
	_ = op.ConnMgoDB()
	err := op.InsertMgoDB(&p1)
	if err != nil {
		fmt.Println("插入错误！")
	} else {
		fmt.Println("插入成功")
	}
}

//Operater MongoDB数据库结构
type Operater struct {
	mgoSession *mgo.Session
	dbname     string
	document   string
}

//Line 业务号结构
type Line struct {
	ID      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`
	Account string        `bson:"account"`
	Fans    int           `bson:"fans"`
	Paas    int           `bson:"paas"`
	Status  int           `bson:"status"`
}

//ConnMgoDB 链接MongoDB数据库
func (operater *Operater) ConnMgoDB() error {
	mgosession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	operater.mgoSession = mgosession
	return nil
}

//InsertMgoDB  插入函数
func (operater *Operater) InsertMgoDB(l *Line) error {
	// p1 := Line{
	// 	// ID:      1,
	// 	Name:    "黄洁",
	// 	Account: "@vfc2617i",
	// 	Fans:    5134,
	// 	Paas:    1,
	// 	Status:  1,
	// }
	// p2 := Line{
	// 	// ID:      2,
	// 	Name:    "朱芷珊",
	// 	Account: "@pbh5563v",
	// 	Fans:    3276,
	// 	Paas:    1,
	// 	Status:  1,
	// }
	// p3 := Line{
	// 	// ID:      3,
	// 	Name:    "李海玲",
	// 	Account: "@jzj5119k",
	// 	Fans:    3620,
	// 	Paas:    1,
	// 	Status:  1,
	// }
	// err := c.Insert(&p1, &p2, &p3)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	collection := operater.mgoSession.DB(operater.dbname).C(operater.document)
	err := collection.Insert(l)
	return err
}

// //GetDataFans 查询函数
// func GetDataFans() {
// 	c := ConnMgoDB()
// 	result := Line{}
// 	err := c.Find(bson.M{"fans": 3276}).One(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("line", result)

// 	lines := make([]Line, 20)
// 	err = c.Find(nil).All(&lines)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(lines)
// }

// //GetDataFansByID 根据ID查询数据
// func GetDataFansByID() {
// 	id := bson.ObjectIdHex("1")
// 	c := ConnMgoDB()
// 	line := &Line{}
// 	err := c.FindId(id).One(line)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(line)
// }

// //UpdateMgoDB 更新函数
// func UpdateMgoDB() {
// 	c := ConnMgoDB()
// 	err := c.Update(bson.M{"fans": 3276}, bson.M{"$set": bson.M{"paas": 2, "status": 0}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// //RemoveMgoDB 删除函数
// func RemoveMgoDB() {
// 	c := ConnMgoDB()
// 	_, err := c.RemoveAll(bson.M{"account": "@jzj5119k"})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
