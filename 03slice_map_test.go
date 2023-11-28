/**
User: cr-mao
Date: 2023/11/28 10:25
Email: crmao@qq.com
Desc: 03slice_map.go
*/
package go_avoid_traps

import (
	"fmt"
	"testing"
)

type Task struct {
	TaskId   int64
	TaskName string
}

func TestSliceMapWrongCase(t *testing.T) {
	var s1 = []Task{
		{
			1,
			"task_1",
		},
		{
			2,
			"task_2",
		},
		{
			3,
			"task_3",
		},
	}

	var m1 = map[int64]*Task{}

	for _, s := range s1 {
		m1[s.TaskId] = &s
	}
	fmt.Println(m1) // map[1:0xc000124078 2:0xc000124078 3:0xc000124078]
	for key, value := range m1 {
		fmt.Println(key, value.TaskId)
		/*
			1 3
			2 3
			3 3
		*/
	}
}

func TestSliceMapRightCase(t *testing.T) {
	var s1 = []Task{
		{
			1,
			"task_1",
		},
		{
			2,
			"task_2",
		},
		{
			3,
			"task_3",
		},
	}

	var m1 = map[int64]*Task{}

	for _, s := range s1 {
		tmp := s
		m1[tmp.TaskId] = &tmp
	}
	fmt.Println(m1) // map[1:0xc000124078 2:0xc000124078 3:0xc000124078]
	for key, value := range m1 {
		fmt.Println(key, value.TaskId)
		/*
			1 1
			2 2
			3 3
		*/
	}
}
