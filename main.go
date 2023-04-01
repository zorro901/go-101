package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2)

	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3) // 5

	i4 := strings.IndexAny("ABC", "ABC")
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5)

	b1 := strings.HasPrefix("ABC", "A") // 検索対象の文字列が指定対象から始まるか
	b2 := strings.HasSuffix("ABC", "C") // 検索対象の文字列が指定対象で終わるか
	fmt.Println("HasPrefix", b1, "HasSuffix", b2)

	b3 := strings.Contains("ABC", "B")      // 検索対象の文字列が含まれるか
	b4 := strings.ContainsAny("ABCDE", "X") // 検索対象の文字列が指定のどこかで含まれるか
	fmt.Println("Contains", b3, "ContainsAny", b4)

	i6 := strings.Count("ABCABC", "B") // 文字列を検索対象として検索した数を返す
	i7 := strings.Count("ABCABC", "")  // 空文字を指定した場合は全部の文字数を返す
	fmt.Println("Count", i6, "CountAny", i7)

	s3 := strings.Repeat("ABC", 4) // 対象文字列を指定した回数だけ繰り返す
	s4 := strings.Repeat("ABC", 0) // ゼロ回は一度も出力しない
	fmt.Println("Repeat", s3, s4)

	s5 := strings.Replace("AAAAAA", "A", "B", 1)  // 対象文字列を指定した文字列に置換
	s6 := strings.Replace("AAAAAA", "A", "B", -1) // -1の場合は全ての置換
	fmt.Println("Replace", s5, s6)

	s7 := strings.Split("A,B,C,D,E", ",") // 分割した文字列を配列として返す
	fmt.Println("Split", s7)
	s8 := strings.SplitAfter("A,B,C,D,E", ",") // セパレータを取り除かない、カンマを付けて分割する
	fmt.Println("SplitAfter", s8)
	s9 := strings.SplitN("A,B,C,D,E", ",", 2) // 分割する最大数を指定する
	fmt.Println("SplitN", s9)
	s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4) // 分割する最大数を指定して、セパレータを取り除かない
	fmt.Println(s10)

	s11 := strings.ToLower("ABC") // 全ての文字を小文字に変換
	s12 := strings.ToLower("E")

	s13 := strings.ToUpper("abc") // 全ての文字を大文字に変換
	s14 := strings.ToUpper("e")
	fmt.Println(s11, s12, s13, s14)

	s15 := strings.TrimSpace("    -    ABC    -    ") // 文字列の前後のスペースを削除
	s16 := strings.TrimSpace("\tABC\r\n")             // 特殊文字も削除される
	s17 := strings.TrimSpace("　　　　ABC　　　　")           // 空白文字を削除
	fmt.Println(s15, s16, s17)

	s18 := strings.Fields("a b c") // 分割した文字列を配列として返す
	fmt.Println(s18)
	fmt.Println(s18[1])
}
