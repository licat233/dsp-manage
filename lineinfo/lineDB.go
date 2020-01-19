package lineinfo

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
	err := operator.ConnMgoDB("127.0.0.1", "dsp", "line")
	if err != nil {
		fmt.Println(err)
		fmt.Println("数据库链接失败")
	}
}

//InsertToMgo 业务号信息注册函数
func InsertToMgo(allLine map[string]string) {
	today := time.Now().Format("2006-01-02")
	wg.Add(len(allLine))
	for name, account := range allLine {
		go func(name string, account string) {
			defer wg.Done()
			_, err := operator.GetData("account", account)
			if err == nil {
				fmt.Println(account, "已存在")
				return
			}
			fansCount := FansCount(account)
			p := mongodb.Line{
				Name:    name,
				Account: account,
				Fans:    fansCount,
				Paas:    0,
				Status:  0,
				Date:    today,
			}
			err = operator.InsertMgoDB(&p)
			if err != nil {
				fmt.Println(account, "注册错误！")
			} else {
				fmt.Println(account, "注册成功！")
			}
		}(name, account)
	}
	wg.Wait()
	fmt.Println("全部注册成功！")
}
