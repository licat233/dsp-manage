package main

import (
	"dspmanage/lineinfo"
)

func main() {
	allMap := map[string]string{"蔡雪娣": "@wfb3891i", "黄妙琼": "@cao9297z"}
	lineinfo.InsertToMgo(allMap)
	// allPopin := []string{"5e0dc0d40fc10c6c215898d4", "5e100cd30fc10c4dff59b175"}
	// popindata.InsertwhiteToMgo(allPopin)

	// name := "first project"
	// account := "self_white_line"
	// campaignid := []string{"5e0dc0d40fc10c6c215898d4", "5e100cd30fc10c4dff59b175"}
	// lineid := []string{"@vfc2617i", "@pbh5563v"}
	// err := project.NewProject(name, account, campaignid, lineid)
	// if err != nil {
	// 	fmt.Println("项目创建失败！")
	// }
}
