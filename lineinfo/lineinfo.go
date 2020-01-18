package lineinfo

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	//line账号
	account string
	//line昵称
	name string
	//全局集合，用于存储line账号
	lineMap = map[string]string{"L蔡雪娣2": "@dtd3861q", "L黄洁": "@vfc2617i", "L朱芷珊1": "@pbh5563v", "L蔡雪娣": "@wfb3891i", "L黄妙琼1": "@cao9297z", "L吴德珊": "@hpp1714r", "L刘娇": "@evu6146b", "L李海玲1": "@jzj5119k", "L何远东": "@wbw6152k", "L黎李结": "@ubu2335t", "L丁思敏": "@kol5958d", "L粱任萍": "@570rreep"}
	//全局集合，用于存储line粉丝数
	fansMap map[string]int
	//全局通道，用于存储line信息
	//lineChan chan string
	//全局通道，用于存储line@粉丝数
	fansChan chan int
	//获取一个随机数
	r *rand.Rand
	//并发控制器，打组
	wg sync.WaitGroup
)

//TotalFans 所有line@账号的粉丝数
func TotalFans(allLine map[string]string) int {
	var sumFans int
	fansChan = make(chan int)
	defer close(fansChan)
	go func() {
		for {
			select {
			case n := <-fansChan:
				sumFans += n
			default:

			}
		}
	}()

	wg.Add(len(lineMap))
	for _, v := range lineMap {
		go Fans(v)
	}
	wg.Wait()

	//time.Sleep(3 *time.Second)

	if sumFans == 0 {
		fmt.Println("查询超时！！")
	}

	return sumFans
}

//常规方式
//func TotalFansCG() int {
//	var sumFans int
//	for _, v := range lineMap {
//		sumFans +=FansCG(v)
//	}
//
//	return sumFans
//}

//GetFans 获取指定line@账号的粉丝数,并发方式
func GetFans(account []string) int {
	var sumFans int
	fansChan = make(chan int)
	defer close(fansChan)
	go func() {
		for {
			select {
			case n := <-fansChan:
				sumFans += n
			default:

			}
		}
	}()

	for _, v := range account {
		go Fans(v)
	}

	if sumFans == 0 {
		fmt.Println("查询超时！！")
	}

	return sumFans
}

//Fans 获取line@账号的粉丝数，不返回，传入通道
func Fans(account string) {
	defer wg.Done()
	account = url.QueryEscape(account)
	var fans int
	webURL := "https://social-plugins.line.me/widget/friend?lineId=" + account + "&count=true&lang=en&origin=https://www." + randString(5) + ".html&title=" + randString(5)
	res, err := http.Get(webURL)
	if err != nil || res.StatusCode != 200 {
		log.Fatal("无法访问：", webURL)
	}
	defer res.Body.Close()

	//加载HTML
	doc, e := goquery.NewDocumentFromReader(res.Body)
	if e != nil {
		log.Fatal(e)
	}
	doc.Find(".btnWrap").Find(".bubble").Find(".num").Each(
		func(i int, s *goquery.Selection) {
			res, ok := s.Attr("title")
			if !ok {
				log.Fatal("查找不到粉丝数量！")
			}
			fans, err = strconv.Atoi(res)
			if err != nil {
				log.Fatal("字符串类型转化错误")
			}
		})

	fansChan <- fans
}

//FansCount @账号的粉丝数，并返回
func FansCount(account string) int {
	defer wg.Done()
	account = url.QueryEscape(account)
	var fans int
	webURL := "https://social-plugins.line.me/widget/friend?lineId=" + account + "&count=true&lang=en&origin=https://www." + randString(5) + ".html&title=" + randString(5)
	res, err := http.Get(webURL)
	if err != nil || res.StatusCode != 200 {
		log.Fatal("无法访问：", webURL)
	}
	defer res.Body.Close()

	//加载HTML
	doc, e := goquery.NewDocumentFromReader(res.Body)
	if e != nil {
		log.Fatal(e)
	}
	doc.Find(".btnWrap").Find(".bubble").Find(".num").Each(
		func(i int, s *goquery.Selection) {
			res, ok := s.Attr("title")
			if !ok {
				log.Fatal("查找不到粉丝数量！")
			}
			fans, err = strconv.Atoi(res)
			if err != nil {
				log.Fatal("字符串类型转化错误")
			}
		})

	return fans
}

//LineInfo 通过通道获取line账号信息【备用】
//func LineInfo() {
//	lineChan = make(chan string)
//	defer close(lineChan)
//	GetAll()
//	time.Sleep(time.Second)
//	for v := range lineChan {
//		fmt.Println(v)
//	}
//}

//GetAll 获取所有的信息
func GetAll() {
	fansMap = make(map[string]int, len(lineMap))
	for n, v := range lineMap {
		GroupString(v, n)
	}
}

//GroupString 查询【业务员】的【line账号】信息
func GroupString(v string, n string) {
	if v != "" {
		s := n + v
		account = url.QueryEscape(v)
		fansMap[s], name = GetLineData(account)
		var build strings.Builder
		build.WriteString(s)
		build.WriteString(" 粉丝数：")
		build.WriteString(strconv.Itoa(fansMap[s]))
		build.WriteString(" 昵称：")
		build.WriteString(name)
		//lineChan <- build.String()
		fmt.Println(build.String())
	} else {
		fmt.Println("账号为空！")
	}
}

//GetLineData 获取line信息(包含粉丝数和昵称）
func GetLineData(account string) (fans int, name string) {
	//请求HTMLhttps://social-plugins.line.me/widget/friend?lineId=%40dtd3861q&count=true&lang=en&origin=https%3A%2F%2Fsocial-plugins.line.me%2Fen%2Fhow_to_install&title=LINE%20Social%20Plugins
	webURL := "https://social-plugins.line.me/widget/friend?lineId=" + account + "&count=true&lang=en&origin=https://www." + randString(5) + ".html&title=" + randString(5)
	fmt.Println(webURL)
	res, err := http.Get(webURL)
	if err != nil || res.StatusCode != 200 {
		log.Fatal("无法访问：", webURL)
	}
	defer res.Body.Close()

	//加载HTML
	doc, e := goquery.NewDocumentFromReader(res.Body)
	if e != nil {
		log.Fatal(e)
	}
	doc.Find(".btnWrap").Find(".bubble").Find(".num").Each(
		func(i int, s *goquery.Selection) {
			res, ok := s.Attr("title")
			if !ok {
				log.Fatal("查找不到粉丝数量！")
			}
			fans, err = strconv.Atoi(res)
			if err != nil {
				log.Fatal("字符串类型转化错误")
			}
		})

	doc.Find(".btnWrap").Find(".btn").Each(
		func(i int, s *goquery.Selection) {
			res, exists := s.Attr("title")
			if !exists {
				log.Fatal("查找不到账号昵称")
			}
			name = getName(res)
		})

	return fans, name
}

// randString 生成随机字符串
func randString(len int) string {
	r = rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65 + 32
		bytes[i] = byte(b)
	}
	return string(bytes)
}

//getName 用正则表达式提取账号昵称
//需要从字符串内提取出：資深護膚老師-Amy
func getName(str string) string {
	as := strings.Index(str, "as") - 1
	// return strings.Trim(str[4:as], " ")
	return str[4:as]
}
