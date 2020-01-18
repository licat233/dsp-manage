package lineinfo

import (
	"dspmanage/mongodb"
	"fmt"
	"time"
)

//InsertToMgo 业务号信息注册函数
func InsertToMgo(allLine map[string]string) {
	op := new(mongodb.Operater)
	today := time.Now().Format("2006-01-02")
	wg.Add(len(allLine))
	for name, account := range allLine {
		go func(name string, account string) {
			fansCount := FansCount(account)
			p := mongodb.Line{
				Name:    name,
				Account: account,
				Fans:    fansCount,
				Paas:    0,
				Status:  0,
				Date:    today,
			}
			err := op.InsertMgoDB(&p)
			if err != nil {
				fmt.Println(account, "插入错误！")
			} else {
				fmt.Println(account, "插入成功")
			}
		}(name, account)
	}
	wg.Wait()
}
