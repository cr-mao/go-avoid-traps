/**
User: cr-mao
Date: 2023/8/30 08:48
Email: crmao@qq.com
Desc: nil.go
*/
package go_avoid_traps

import (
	"fmt"
	"testing"
)

/*
it's non-empty
it's empty
it's not empty
*/

func TestNil(t *testing.T) {
	var s *string
	test(s)
	res := testResult()
	if res == nil {
		fmt.Println("it's empty")
	} else {
		fmt.Println("it's not empty")
	}

	res2 := testResult2()
	if res2 == nil {
		fmt.Println("it's empty")
	} else {
		fmt.Println("it's not empty")
	}
}

func testResult() *string {
	var s *string
	return s
}
func testResult2() interface{} {
	var s *string
	return s
}

func test(s interface{}) {
	if s == nil {
		fmt.Println("it's empty")
	} else {
		fmt.Println("it's non-empty")
	}
}
