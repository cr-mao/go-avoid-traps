/**
User: cr-mao
Date: 2023/8/26 14:01
Email: crmao@qq.com
Desc: rand_num_test.go
*/
package go_avoid_traps

import (
	"math/big"
	mathRand "math/rand"
	"testing"
	"time"

	"crypto/rand"
)

// 坑, 可能会留在一个区间，达不到你想要的效果, 抽奖程序，可能连续抽到好的，或者差的，工作中有遇到过.
func Test_RandNumByMathRand(t *testing.T) {
	mathRand.Seed(time.Now().UnixNano())
	var result = make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		result = append(result, mathRand.Intn(1000))
	}
	t.Log(result)
}

// 真随机，就是性能慢点,根据场景选择
func realRandNum(num int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(num))
	return n.Int64()
}

func Test_CryptoRand(t *testing.T) {
	var result = make([]int64, 0, 1000)
	for i := 0; i < 1000; i++ {
		result = append(result, realRandNum(1000))
	}
	t.Log(result)
}
