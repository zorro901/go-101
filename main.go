package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return l[i].Name < l[j].Name
	} else {
		return l[i].Value < l[j].Value
	}
}
func main() {

	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	s := []string{"a", "z", "j"}

	fmt.Println(i, s)

	sort.Ints(i)    // [5 3 2 4 5 6 4 8 9 8 7 10] [a z j]
	sort.Strings(s) // [a j z]

	fmt.Println(i, s)

	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	fmt.Println(el)

	// 名前を昇順にソート
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	fmt.Println("名前を昇順にソート")
	fmt.Println(el)
	// 値順にソート、名前順は上書きされる
	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println("---------")
	fmt.Println(el)
	fmt.Println("---------")

	// 名前順を維持して値順にソート
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println("---------")
	fmt.Println(el)
	fmt.Println("---------")

	m := map[string]int{"ada": 1, "hoge": 4, "basha": 3, "poeni": 3}

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	//sort.SliceStable(lt, func(i, j int) bool { return lt[i].Name < lt[j].Name })
	//fmt.Println(lt)

	sort.Sort(lt)

	fmt.Println(lt)

	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)

}
