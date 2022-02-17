/**
 * @Author: jackie
 * @Descripttion: 10进制转换n进制转换
 * @version:
 * @Date: 2022-02-16 14:33:26
 * @LastEditors: jackie
 * @LastEditTime: 2022-02-17 09:50:40
 */
package main

import (
	"math"
)

type BaseN struct {
}

// baseMap 长度标识进制，其中字符标识n进制数
var baseMap []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// AddN n进制numBaseNum加十进制n
// 1. n进制numBaseNum通过map下标转换成十进制，然后加a在转换为n进制
// 2. n进制numBaseNum直接加n,n进制进位
func (bn BaseN) AddN(numBaseNum string, n int) string {
	return bn.ConvertToBaseN(bn.BaseNToTen(numBaseNum) + n)
}

// ConvertToBaseN num转n进制表示
func (bn BaseN) ConvertToBaseN(num int) string {
	n := len(baseMap)
	ans := ""
	flag := false
	if num < 0 {
		flag = true
		num = -num
	}
	for ; num >= n; num = num / n {
		ans = baseMap[num%n] + ans
	}

	ans = baseMap[num] + ans
	if flag {
		ans = "-" + ans
	}
	return ans
}

// BaseNToTen n进制转10进制
func (bn BaseN) BaseNToTen(numBaseNum string) int {
	var ans float64
	for k, _ := range numBaseNum {
		tmp := bn.getKey((string)(numBaseNum[len(numBaseNum)-k-1]))
		if k == 0 { // 最后一位只加
			ans += tmp
		} else {
			ans += tmp * math.Pow(float64(len(baseMap)), float64(len(numBaseNum)-1))
		}
	}

	return int(ans)
}

// getKey 通过字符获取下标key
func (bn BaseN) getKey(m string) float64 {
	for k, v := range baseMap {
		if v == m {
			return float64(k)
		}
	}
	return 0
}

// SetBase 自定义进制
func SetBase(nbase []string) {
	baseMap = nbase
}
