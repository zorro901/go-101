package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("A", "ABC") // 基本形、簡易
	fmt.Println(match)

	re1, _ := regexp.Compile("A") // 正規表現のパターンを予め生成しておく、動的に正規表現が変化する場合に対応できる
	match = re1.MatchString("ABC")
	fmt.Println("Compile:", match)

	re2 := regexp.MustCompile("A") // コンパイルエラーが発生する場合、直接ランタイムエラーを発生させる
	match = re2.MatchString("ABC")
	fmt.Println(match)

	// エスケープの書き方
	//regexp.MustCompile("\\d")
	//regexp.MustCompile(`\d`)

	re3 := regexp.MustCompile(`(?im)abc`) // i 大文字小文字を区別しない
	match = re3.MatchString("ABC")
	fmt.Println("MustCompile", match)

	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match)
	match = re4.MatchString("  ABC  ")
	fmt.Println(match)

	/*
		re := regexp.MustCompile("ab")
		re.MatchStrings("abc")
		//true
	*/

	re5 := regexp.MustCompile(".")
	match = re5.MatchString("ABC")
	fmt.Println(match)
	match = re5.MatchString("\n")
	fmt.Println(match)

	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab"))
	fmt.Println(re6.MatchString("a"))
	fmt.Println(re6.MatchString("aaaabbbbbbbb"))
	fmt.Println(re6.MatchString("b"))

	re7 := regexp.MustCompile("A+?A+?X")
	fmt.Println(re7.MatchString("AAX"))
	fmt.Println(re7.MatchString("AAAAAAXXXX"))

	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y"))

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re9.MatchString("ABC"))
	fmt.Println(re9.MatchString("abcdefg"))

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re10.MatchString("ABC"))
	fmt.Println(re10.MatchString("あ"))

	re1 = regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re1.MatchString("abcxyz"))
	fmt.Println(re1.MatchString("ABCXYZ"))
	fmt.Println(re1.MatchString("ABCxyz"))
	fmt.Println(re1.MatchString("ABCabc"))

	fs1 := re1.FindString("AAAAABCXYZZZZ")
	fmt.Println("FindString", fs1)
	fs2 := re1.FindAllString("ABCXYZABCXYZ", -1) // -1の指定でマッチした全ての文字列を取得する
	fmt.Println("FindAllString", fs2)

	fmt.Println(re1.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1)) // [ASHVJV<H KNJBJVK ]
	re1 = regexp.MustCompile(`\s+`)                           // 1つ以上の連続した空白文字にマッチする
	fmt.Println(re1.Split("aaaaaaaaaa     bbbbbbb  cccccc", -1))

	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ",")) // AAA,BBB,CCC
	re1 = regexp.MustCompile(`、|。`)
	fmt.Println(re1.ReplaceAllString("私は、Golangを使用する、プログラマー。", ""))

	re1 = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
	0123-456-7889
	111-222-333
	556-787-899
	`

	ms := re1.FindAllStringSubmatch(s, -1)
	fmt.Println("FindAllStringSubmatch", ms) // 先頭にマッチした文字列の全体、それ以降の要素はサブマッチした要素を格納したスライス
	for _, v := range ms {
		fmt.Println(v)
	}

	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "$3-$2-$1"))
	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "あ-い-う"))

}
