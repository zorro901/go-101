package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

func f[T Stringer](xs []T) []string {
	result := []string{}
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	fmt.Println(f([]MyInt{1, 2, 3}))
}
