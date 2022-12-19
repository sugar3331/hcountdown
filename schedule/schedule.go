package schedule

import (
	"fmt"
	"time"
)

var schMap = make(map[string]string)

var mm = ""
var ml = ""
var mt = 999

func init() {

}

//创建日程
func CreatedSch(s string, t time.Time) {
	ns := t.Format("2006年1月2日")
	schMap[s] = ns
}

//删除日程
func DeletedSch(s string) {
	delete(schMap, s)
}

//查询日程
func QuerySch() {
	if len(schMap) == 0 {
		fmt.Println("日程为空")
	} else {
		for country := range schMap {
			fmt.Println("安排日程是：", country, "日期为:", schMap[country])
		}
	}
	fmt.Println(len(schMap))
}

//判断最近的日程
func JudgeSch(t time.Time) {
	s1 := t.Format("2006年1月2日")
	t1, _ := time.Parse("2006年1月2日", s1)
	for c := range schMap {
		t2, err := time.Parse("2006年1月2日", schMap[c])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		d1 := t2.Sub(t)
		d2 := time.Duration.Hours(d1)
		dd1 := int(d2)
		d3 := int(d2 / 24)
		if t1 == t2 {
			fmt.Println("今天是规划中的日程 ： ", c)
			return
		} else if dd1 < 0 {
			continue
		} else {

			if d3 < mt {
				ml = schMap[c]
				mm = c
				mt = d3
			}
		}
	}
	fmt.Println("距离现在最近的日程是： ", mm, "日期为： ", "还有 ", mt+1, "天")
}
