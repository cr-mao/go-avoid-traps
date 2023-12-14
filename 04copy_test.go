/**
User: cr-mao
Date: 2023/12/15 07:34
Email: crmao@qq.com
Desc: 深浅copy
*/
package go_avoid_traps

import (
	"fmt"
	"testing"
)

type Player struct {
	UserId int64 `json:"user_id"`
	store  map[int32]int32
}
type Players struct {
	inner map[int64]*Player
}

func TestDeepCopy(t *testing.T) {
	var group = Players{
		inner: map[int64]*Player{
			1: {
				UserId: 1,
				store:  map[int32]int32{20: 20},
			},
		},
	}
	copyPlayer := *group.inner[1]
	// 得重新复制， 有指针得还不行哦。
	copyPlayer.store = make(map[int32]int32)
	for key, v := range group.inner[1].store {
		copyPlayer.store[key] = v
	}
	copyPlayer.store[21] = 21
	fmt.Println(group.inner[1].store) // 没改变 20：20
}

func TestShallowCopy(t *testing.T) {
	var group = Players{
		inner: map[int64]*Player{
			1: {
				UserId: 1,
				store:  map[int32]int32{20: 20},
			},
		},
	}
	copyPlayer := *group.inner[1]
	copyPlayer.store[21] = 21
	fmt.Println(group.inner[1].store) // 没改变 20：20 ,21:21
}
