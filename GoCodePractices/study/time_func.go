package main

import (
	"fmt"
	"time"
)

const (
	DATE = "2006-01-02"
	TIME = "2006-01-02 15:04:40.000" // 这只是一种用于显示时间时所展示的格式，即冒号可以替换为`-`等
)

func main() {
	t0 := time.Now()                                   // 获取本地时间 time.Time类型 包含本地时区信息
	fmt.Println("Original timestamp:", t0.UnixMilli()) // 时间戳
	utcTime := t0.UTC()                                // 获取t0时间对应的UTC时间
	fmt.Println("UTC timestamp:", utcTime)
	// 获取本地时间的时区
	zoneName, offset := t0.Zone()
	fmt.Println("时区偏移量(秒)：", offset)
	// 获取时区的名称
	fmt.Println("时区名称：", zoneName) // CST

	// 暂停时间
	time.Sleep(51 * time.Millisecond)
	//获取时间差
	t1 := time.Now()
	diff := t1.Sub(t0)
	fmt.Println(diff.Milliseconds())

	//Since
	fmt.Println(time.Since(t0).Milliseconds()) // 等同于15 16行

	// Time + Duration = Time
	d := 5 * time.Second
	t2 := t0.Add(d)
	fmt.Println(t2.Unix())

	// 年月日  时分秒
	//fmt.Println(t0.Format(DATE))
	s := t0.Format(TIME)
	fmt.Println(s)
	//
	t3, _ := time.Parse(TIME, s)               // 将s时间字符串以TIME这个格式解析为time.Time
	fmt.Println("Parsed timestamp", t3.Unix()) // 这里的时间戳打印值与之前的t0打印不同，因为存在时间的损失，具体是因为有毫秒的损失
	// 这里替代使用UnixMilli() 获取到精确到毫秒的时间戳,能以避免时间损失
	fmt.Println("Parsed timestamp", t3.UnixMilli())

	// Time zone example
	location, _ := time.LoadLocation("America/New_York") // 加载特定时区
	l, _ := time.LoadLocation("Europe/London")
	newYorkTime := t0.In(location) // 将本地时间转换为特定时区的时间
	londonTime := t0.In(l)
	fmt.Println("New York Time:", newYorkTime)
	fmt.Println("London Time:", londonTime)
	fS, _ := time.ParseInLocation(TIME, s, l) // 将s时间字符串以TIME这个格式解析为time.Time，且设置时区为London
	fmt.Println("f_s:", fS)
}
