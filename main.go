package main

import (
	"dspmanage/project"
	"fmt"
)

func main() {
	// allMap := map[string]string{"黄洁": "@vfc2617i", "朱芷珊": "@pbh5563v"}
	// lineinfo.InsertToMgo(allMap)
	// allPopin := []string{"5e0dc0d40fc10c6c215898d4", "5e100cd30fc10c4dff59b175"}
	// popindata.InsertwhiteToMgo(allPopin)

	name := "first project"
	account := "self_white_line"
	campaignid := []string{"5e0dc0d40fc10c6c215898d4", "5e100cd30fc10c4dff59b175"}
	lineid := []string{"@vfc2617i", "@pbh5563v"}
	err := project.NewProject(name, account, campaignid, lineid)
	if err != nil {
		fmt.Println("项目创建失败！")
	}
}
