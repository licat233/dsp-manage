package project

import (
	"dspmanage/mongodb"
	"fmt"
)

var (
	//活动ID
	campaignID []string
	//line账号ID
	lineID []string
	//实例化数据库
	operator = new(mongodb.Operater)
)

func init() {
	err := operator.ConnMgoDB("127.0.0.1", "dsp", "project")
	if err != nil {
		fmt.Println(err)
		fmt.Println("数据库链接失败")
	}
}

//NewProject 新建项目
func NewProject(name string, account string, campaignID []string, lineID []string) error {
	p := mongodb.Project{
		Name:       name,
		Campaignid: campaignID,
		Lineid:     lineID,
		Account:    account,
	}
	err := operator.InsertProject(&p)
	if err != nil {
		return err
	}
	return nil
}
