package main

import (
	"fmt"
	"gopacong/lineinfo"
	"time"
)

func data() {
	//a := popin.WhiteMoney()
	//fmt.Println(a)
	//b := popin.WhiteLineMoney()
	//fmt.Println(b)
	start := time.Now()
	f := lineinfo.TotalFans()
	fmt.Println("总粉丝数:", f)
	fmt.Println("用时:", time.Since(start))
}
