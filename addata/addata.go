package addata

import (
	"dspmanage/lineinfo"
	"dspmanage/mongodb"
	"fmt"
	"gopacong/popindata"

	"gopkg.in/mgo.v2/bson"
)

var (
	//fansTotal 相关line账号的总粉丝数,标记
	fansTotal int
	//nowFansTotal 相关line账号的总粉丝数,实时
	nowFansTotal int
	//moneyTotal 相关popin活动的总消耗,标记
	moneyTotal int
	//nowMoneyTotal 相关popin活动的总消耗,实时
	nowMoneyTotal int
	//实例化数据库
	operator           = new(mongodb.Operater)
	whiteID     string = "Self_White"
	whiteLineID string = "self_white_line"
)

func init() {
	err := operator.ConnMgoDB("127.0.0.1", "dsp", "popin")
	if err != nil {
		fmt.Println(err)
		fmt.Println("数据库链接失败")
	}
}

//ProjectCPA 项目的CPA
func ProjectCPA(projectid bson.ObjectId) {
	project, err := operator.GetProjectByID(projectid)
	if err != nil {
		fmt.Println("查找不到项目！")
	}
	if project.Account == whiteID {

	}
	nowFansTotal = lineinfo.GetFans(project.Lineid)
	nowMoneyTotal = popindata.WhiteMoneyDIY(project.Campaignid)

}
