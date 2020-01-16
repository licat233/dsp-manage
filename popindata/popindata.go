package popindata

import (
	"encoding/json"
	"fmt"
	"gopacong/common"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var (
	//POPin平台Self_White账户的cookie
	cookieWhite string = "PHPSESSID=oc9kkk64ro3t649usg7g2v5h8a; __stripe_mid=94a1c653-251f-493c-a911-0c2afd3f97f2; autoLoginEnable=true; __zlcmid=wCipPAqVpFkh8H; laravel_session=eyJpdiI6IjJuMzBJckZMZVFQSFJEY29mNEZcL0R3PT0iLCJ2YWx1ZSI6ImpNVmlWMjdxZkpBbm9yZjdLUXNqck1UTXJpUm9yN1FKWXZPQmxqdU9MSUg1M2dBN0JJNmI5VFNSVXlNRlM1S2ciLCJtYWMiOiJiZWI1M2E4M2E4ZWU2YjZiNmViYmIyOTAxZDJhYzFmZDAzY2ExN2UyNzNiZGZjN2QzNDc5MjBmZDY0OTg2ZjYwIn0%3D; __stripe_sid=35c49d73-e5b3-440d-8ffd-2337b33662a2; ticket=94c85e7803"

	//请求Html的self_Whit账户的cookie
	cookieWhiteHTML string = "PHPSESSID=jmg8f2upol286c2k22porrbcmb; laravel_session=eyJpdiI6Inc2UkJ4cFJoSmFWandmTEZ6NXRiTUE9PSIsInZhbHVlIjoiWHo4NXVnTkRqZGFTbjFOaG1Db2FrcDZhMEU5K1lYSVFUWlVqTXJSWU9nem15a1FhZmk5eDgxcHMwZjFIS2J4RCIsIm1hYyI6IjAxN2IwMmQ4MmE0MmY3YjNlNDE0YjM1Mjg4NDYzOGRhOTM5Zjg2ODhkNmEwOGZkMDg4YmVmMWY3ZGYxNWE2ZGMifQ%3D%3D; __stripe_mid=56c4633c-c270-4d6f-bc4f-05152d7b6e12; __stripe_sid=316113f1-bfc1-4b1c-b2a9-712d93e80172; autoLoginEnable=true; ticket=94c85e7803; __zlcmid=wCipSWW4SZeYGz"

	//POPin平台Self_White_line账户的cookie
	cookieWhiteLine string = "PHPSESSID=slm4a6o7jl3ql760rbln1k0bur; laravel_session=eyJpdiI6IjBSTUg2dytrOVAwWGtCdXJUem8zWVE9PSIsInZhbHVlIjoiY0k1TkFBSG9HYWJUNEE1dVZwa0lZamhSWGlxQzZYalFmQTlLRSt4eGtFa3g3Q2hsbkNFT1I3VXRwSUU1RnBuXC8iLCJtYWMiOiJkNjg3MGIwMmFkMzFlMWU2NmE1YzRmZDg4ZGI2ZDFiZTQ4NDJhNjVhOGFlYjQzMjYyYzdjYTdiYWM0OGU0ODc4In0%3D; __stripe_mid=42533a17-6c5b-436b-bda3-fe036aebe216; __stripe_sid=2cded61c-b576-4de8-a6a2-bca341be4a40; autoLoginEnable=true; ticket=22e9205fc8; __zlcmid=wCipRT9bUpBgtF"

	//请求Html的self_White_line账户的cookie
	cookieWhiteLineHTML string = "PHPSESSID=slm4a6o7jl3ql760rbln1k0bur; laravel_session=eyJpdiI6InBSYm1ueWtrYnp1Yk0wcjEyTmxOYnc9PSIsInZhbHVlIjoiNnVEYk5JNHJ1TDZPUkJPYU1wNG5NcVNiSEtLU1huTmVVMDFiUzNZdmtRTnh6TTZUUmxleFwvdHl5c2JYQnlaWVEiLCJtYWMiOiJjZWQ2ZTBlODc3NTA2MzFkOTMwMmVhNjdjMjI1YWJhMDUwNWMyNGFiZWE2MGUyNjA1OWE0MzQ4NDBmZmNlZGQ5In0%3D; __stripe_mid=42533a17-6c5b-436b-bda3-fe036aebe216; __stripe_sid=2cded61c-b576-4de8-a6a2-bca341be4a40; autoLoginEnable=true; ticket=22e9205fc8; __zlcmid=wCipRT9bUpBgtF"

	whiteID     string = "Self_White"
	whiteLineID string = "self_white_line"
)

type (
	popinJSONObj struct {
		TotalMoney string `json:"total_charge"`
	}
)

//WhiteMoney 获取whiteID的ON活动的总消耗
func WhiteMoney() int {
	idRes := GetCampaignID(cookieWhiteHTML, whiteID)
	var money int
	for i, v := range idRes {
		j := GetMoney(cookieWhite, whiteID, v)
		fmt.Println(i, ":", j)
		money += j
	}
	return money
}

//WhiteLineMoney 获取whiteLineID的ON活动的总消耗
func WhiteLineMoney() int {
	idRes := GetCampaignID(cookieWhiteLineHTML, whiteLineID)
	var money int
	for i, v := range idRes {
		j := GetMoney(cookieWhiteLine, whiteLineID, v)
		fmt.Println(i, ":", j)
		money += j
	}
	return money
}

//WhiteMoneyDIY 获取whiteID手动指定活动ID的总消耗
func WhiteMoneyDIY(idSlice []string) int {
	var money int
	for i, v := range idSlice {
		j := GetMoney(cookieWhite, whiteID, v)
		fmt.Println(i, ":", j)
		money += j
	}
	return money
}

//WhiteLineMoneyDIY 获取whiteLineID手动指定活动ID的总消耗
func WhiteLineMoneyDIY(idSlice []string) int {
	var money int
	for i, v := range idSlice {
		j := GetMoney(cookieWhiteLine, whiteLineID, v)
		fmt.Println(i, ":", j)
		money += j
	}
	return money
}

//GetMoney 爬取所需要的数据（消耗额）
func GetMoney(cookie, accountID, campaignID string) int {
	url := "https://dashboard.popin.cc/discovery/accounts-tw/index.php/" + accountID + "/c/" + campaignID + "/getDiscoveryReports"
	referer := "https://dashboard.popin.cc/discovery/accounts-tw/index.php/" + accountID + "/manageAgency?subPage=/" + accountID + "/campaigns/listCampaign"

	client := &http.Client{}
	var resp *http.Response
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("http.NewRequest Failed:", err)
		//os.Exit(1)
	}

	req.Header.Add("Host", "dashboard.popin.cc")
	req.Header.Add("User-Agent", common.GetRandomUserAgent())
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", referer)
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")

	//发送请求数据之服务器，并获取响应数据
	resp, err = client.Do(req)
	//time.Sleep(5 * time.Second)

	if err != nil {
		log.Fatal("Client Do Failed:", err)
		//os.Exit(1)
	}
	if resp.StatusCode != 200 {
		log.Fatal("status code error: ", resp.StatusCode, resp.Status)
		//os.Exit(1)
	}

	// 打印所有请求头信息
	//for k, v := range req.Header {
	//	fmt.Println("请求头信息 :", k, "=", v)
	//}

	// 打印所有响应头信息
	//for k, v := range resp.Header {
	//	fmt.Println("响应头信息 :", k, "=", v)
	//}

	//读取JSON文件内容 返回字节切片
	bytesSlice, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("ioutil.ReadAll failed:", err)
		//os.Exit(1)
	}

	//开始处理信息
	var popJSON popinJSONObj
	//解码json数据，将字节切片映射到指定结构上
	e := json.Unmarshal(bytesSlice, &popJSON)
	if e != nil {
		log.Fatal("json.Unmarshal:", e)
		//os.Exit(1)
	}

	//fmt.Println(popJSON.TotalMoney)

	moneystring := strings.Replace(popJSON.TotalMoney, ",", "", -1)
	moneystring = moneystring[:len(moneystring)-3]

	moneyint, err := strconv.Atoi(moneystring)
	if err != nil {
		log.Fatal("strconv.Atoi failed:", err)
		//os.Exit(1)
	}

	return moneyint
}

//GetCampaignID 获取活动的ID
func GetCampaignID(cookie, accountID string) []string {
	url := "https://dashboard.popin.cc/discovery/accounts-tw/index.php/" + accountID + "/campaigns/listCampaign"
	referer := "https://dashboard.popin.cc/discovery/accounts-tw/index.php/" + accountID + "/manageAgency?subPage=/" + accountID + "/campaigns/listCampaign"
	client := &http.Client{}
	var resp *http.Response
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("http.NewRequest Failed:", err)
		//os.Exit(1)
	}

	req.Header.Add("Host", "dashboard.popin.cc")
	req.Header.Add("User-Agent", common.GetRandomUserAgent())
	req.Header.Add("Accept", "text/html, */*; q=0.01")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", referer)
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Pragma", "no-")
	req.Header.Add("Cache-Control", "no-cache")

	//发送请求数据之服务器，并获取响应数据
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal("Client Do Failed:", err)
		//os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("status code error:", resp.StatusCode, resp.Status)
		//os.Exit(1)
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll failed:", err)
		//os.Exit(1)
	}
	content := string(byt)
	//fmt.Println(content)

	//正则表达式
	regID := regexp.MustCompile(`data-campaignid="[a-zA-Z0-9]+"`)
	regStatus := regexp.MustCompile(`campaign-status\sstring">[A-Z]+</td>`)
	idSlice := regID.FindAllString(content, -1)
	statusSlice := regStatus.FindAllString(content, -1)
	//var statusChan chan int
	statusChan := make(chan int)
	defer close(statusChan)
	var campaignID []string

	go func() {
		for {
			select {
			case i := <-statusChan:
				v := idSlice[i][17 : len(idSlice[i])-1] //已获取ON的活动ID
				campaignID = append(campaignID, v)
			default:

			}
		}
	}()

	for i, va := range statusSlice {
		if strings.Index(va, "ON") != -1 {
			//va = va[24 : len(va)-5]
			statusChan <- i
		}
	}

	return campaignID
}
