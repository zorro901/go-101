package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	t2 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay())
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())

	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	// 2分15秒後の時間を生成
	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t3)

	// t3-t(2分15秒後の時間-現在時刻)
	d2 := t3.Sub(t)
	fmt.Println(d2)

	fmt.Println(t.Before(t3))
	fmt.Println(t.After(t3))
	fmt.Println(t3.Before(t))
	fmt.Println(t3.After(t))

	time.Sleep(5 * time.Second)
	fmt.Println("5秒停止後")
}
