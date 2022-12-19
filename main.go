package main

import (
	"github.com/sugar3331/hcountdown/reptile"
	"github.com/sugar3331/hcountdown/schedule"
	"time"
)

func main() {
	//现在的时间
	nt := time.Now()
	//显示距离最近的节日
	reptile.GetNear(nt)

	rm := "freetime"                                //添加的日程名字
	rt, _ := time.Parse("2006年1月2日", "2022年12月18日") //添加的日程时间
	schedule.CreatedSch(rm, rt)

	rm2 := "freetime2"                               //添加的日程名字
	rt2, _ := time.Parse("2006年1月2日", "2022年12月25日") //添加的日程时间
	schedule.CreatedSch(rm2, rt2)

	rm3 := "睡觉"                                      //添加的日程名字
	rt3, _ := time.Parse("2006年1月2日", "2022年12月14日") //添加的日程时间
	schedule.CreatedSch(rm3, rt3)

	schedule.QuerySch()
	schedule.DeletedSch("睡觉")
	schedule.QuerySch()

	schedule.JudgeSch(time.Now())
}
