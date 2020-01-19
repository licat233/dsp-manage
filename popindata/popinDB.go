package popindata

import (
	"dspmanage/mongodb"
	"fmt"
	"time"
)

var (
	//实例化数据库
	operator = new(mongodb.Operater)
)

func init() {
	err := operator.ConnMgoDB("127.0.0.1", "dsp", "popin")
	if err != nil {
		fmt.Println(err)
		fmt.Println("数据库链接失败")
	}
}

//InsertwhiteToMgo white活动注册函数
func InsertwhiteToMgo(allPopin []string) {
	today := time.Now().Format("2006-01-02")
	wg.Add(len(allPopin))
	for _, campaignid := range allPopin {
		go func(campaignid string) {
			defer wg.Done()
			_, err := operator.GetData("campaignID", campaignid)
			if err == nil {
				fmt.Println(campaignid, "已存在")
				return
			}
			money := GetWhiteMoney(campaignid)
			p := mongodb.Popin{
				CampaignID: campaignid,
				Totalmoney: money,
				Status:     0,
				Date:       today,
			}
			err = operator.InsertMgoDBP(&p)
			if err != nil {
				fmt.Println(campaignid, "注册错误！")
			} else {
				fmt.Println(campaignid, "注册成功！")
			}
		}(campaignid)
	}
	wg.Wait()
	fmt.Println("全部注册成功！")
}

//InserttodayToMgo whitelinetoday活动注册函数
func InserttodayToMgo(allPopin []string) {
	today := time.Now().Format("2006-01-02")
	wg.Add(len(allPopin))
	for _, campaignid := range allPopin {
		go func(campaignid string) {
			defer wg.Done()
			_, err := operator.GetData("campaignID", campaignid)
			if err == nil {
				fmt.Println(campaignid, "已存在")
				return
			}
			money := GetLineMoney(campaignid)
			p := mongodb.Popin{
				CampaignID: campaignid,
				Totalmoney: money,
				Status:     0,
				Date:       today,
			}
			err = operator.InsertMgoDBP(&p)
			if err != nil {
				fmt.Println(campaignid, "注册错误！")
			} else {
				fmt.Println(campaignid, "注册成功！")
			}
		}(campaignid)
	}
	wg.Wait()
	fmt.Println("全部注册成功！")
}
