package reptile

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var ctf map[string]string

func init() {
	t := time.Now()
	year := t.Year()
	yearStr := strconv.Itoa(year)

	url := crawl(yearStr)
	ctf = deepSource(url, yearStr)
}

func GetNear(t time.Time) (str string, n int) {
	year := t.Year()
	str, n = judge(t, year, ctf)
	return str, n
}

//先在官方网站中爬取全部数据,返回需要的url
func crawl(yearStr string) string {
	url := ""
	resp, err := http.Get("http://sousuo.gov.cn/s.htm?t=govall&advance=false&n=&timetype=&mintime=&maxtime=&sort=&q=%E8%8A%82%E5%81%87%E6%97%A5%E5%AE%89%E6%8E%92")
	if err != nil {
		fmt.Println("初始页面加载失败", err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return err.Error()
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取数据失败", err.Error())
	}

	//在再找的数据中筛选出当年的节假日发布url
	re := regexp.MustCompile(`<a href="(.+)" target="_blank">(国务院办公厅关于` + yearStr + `年部分<em>节假日安排</em>的通知)</a>`)
	matches := re.FindAllSubmatch(all, -1)
	for _, m := range matches {
		//fmt.Printf("url : %s, word : %s\n", m[1], m[2])
		if m[1] != nil {
			url = string(m[1])
			break
		}
	}
	return url
}

//在最终的需要网页中爬取信息，并存入map集合中
func deepSource(url, yearStr string) map[string]string {
	//var ctf map[string]string
	ctf = make(map[string]string)
	resp2, err := http.Get(url)
	if err != nil {
		fmt.Println("节日页面加载失败", err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp2.Body)
	if resp2.StatusCode != http.StatusOK {
		fmt.Println("Error : status code2",
			resp2.StatusCode)
	}

	all2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println("获取数据失败2", err.Error())
	}

	//清洗出最终数据并存入map集合中
	re2 := regexp.MustCompile(`、(.+)：.+([1-9]月.+)至.+天。`)
	matches2 := re2.FindAllSubmatch(all2, -1)
	for _, m := range matches2 {
		str := string(m[2])
		m[2] = []byte(yearStr + "年" + str)
		ctf[string(m[1])] = string(m[2])
	}
	return ctf
}

//逻辑判断距离最近节假日还有多少天
func judge(t time.Time, year int, ctf map[string]string) (str string, n int) {
	mlt := 367
	mlf := "元旦"
	t5, _ := time.Parse("2006年1月2日", "2022年9月11日") //判定时间

	lll := t5.Sub(t)
	if lll <= 0 {
		s4 := strconv.Itoa(year+1) + "年1月1日"
		t6, _ := time.Parse("2006年1月2日", s4)
		d2 := t6.Sub(t)
		tt1 := time.Duration.Hours(d2)
		ls1 := int(tt1 / 24)
		str = mlf
		n = ls1 + 1
	} else {
		for kl := range ctf {
			s4 := ctf[kl]
			t6, _ := time.Parse("2006年1月2日", s4)
			d2 := t6.Sub(t)
			if d2 < 0 {
				continue
			} else {
				tt1 := time.Duration.Hours(d2)
				ls1 := int(tt1 / 24)
				if ls1 < mlt {
					mlt = ls1
					mlf = kl
				}
			}
		}
		str = mlf
		n = mlt + 1
	}
	return str, n
}
