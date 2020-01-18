package mongodb

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Line 业务号结构
type Line struct {
	// ID      bson.ObjectId `bson:"_id"`
	Name    string `bson:"name"`
	Account string `bson:"account"`
	Fans    int    `bson:"fans"`
	Paas    int    `bson:"paas"`
	Status  int    `bson:"status"`
	Date    string `bson:"date"`
}

func demo() {
	p1 := Line{
		// ID:      1,
		Name:    "黄洁",
		Account: "@vfc2617i",
		Fans:    5134,
		Paas:    1,
		Status:  1,
	}
	p2 := Line{
		// ID:      1,
		Name:    "朱芷珊",
		Account: "@pbh5563v",
		Fans:    3276,
		Paas:    1,
		Status:  1,
	}
	p := []Line{p1, p2}
	op := new(Operater)
	e := op.ConnMgoDB("127.0.0.1", "dsp", "line")
	if e != nil {
		fmt.Println("链接失败")
	} else {
		fmt.Println("链接成功")
	}

	for _, v := range p {
		err := op.InsertMgoDB(&v)
		if err != nil {
			fmt.Println("插入错误！")
		} else {
			fmt.Println("插入成功")
		}
	}

	result, err := op.GetData("account", "@pbh5563v")
	if err != nil {
		fmt.Println("查询失败!")
	}
	fmt.Printf("%v\n", result)
	id := bson.ObjectIdHex("5e211de83ca65a5e795bfd8b")
	res, err := op.GetDataByID(id)
	if err != nil {
		fmt.Println("查询失败!")
	}
	fmt.Printf("%v\n", res)
	err = op.UpdateFans("@pbh5563v", 3333)
	if err != nil {
		fmt.Println("更新失败")
	}
	err = op.RemoveAcc("@pbh5563v")
	if err != nil {
		fmt.Println("删除失败")
	}
}

//Operater MongoDB数据库结构
type Operater struct {
	mgoSession *mgo.Session
	database   *mgo.Database
	collection *mgo.Collection
}

//ConnMgoDB 链接MongoDB数据库
func (operater *Operater) ConnMgoDB(host string, dbname string, colle string) error {
	mgosession, err := mgo.Dial(host)
	if err != nil {
		fmt.Println(err)
		return err
	}
	mgosession.SetMode(mgo.Monotonic, true)
	//选择集合
	operater.collection = mgosession.DB(dbname).C(colle)
	return nil
}

//InsertMgoDB  插入函数
func (operater *Operater) InsertMgoDB(p *Line) error {
	err := operater.collection.Insert(p)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//GetData 查询函数传入值类型string
func (operater *Operater) GetData(key string, value string) (Line, error) {
	result := Line{}
	err := operater.collection.Find(bson.M{key: value}).One(&result)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, nil
}

//GetDataByID 根据ID查询数据
func (operater *Operater) GetDataByID(id bson.ObjectId) (Line, error) {
	line := Line{}
	err := operater.collection.FindId(id).One(&line) //一定要传入line的地址，即加&标识符
	if err != nil {
		fmt.Println(err)
		return line, err
	}
	return line, nil
}

//UpdateFans 更新函数,根据账号account修改粉丝数fans
func (operater *Operater) UpdateFans(account string, fans int) error {
	err := operater.collection.Update(bson.M{"account": account}, bson.M{"$set": bson.M{"fans": fans}})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//RemoveAcc 删除函数，传入账号删除指定账号
func (operater *Operater) RemoveAcc(account string) error {
	_, err := operater.collection.RemoveAll(bson.M{"account": account})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
