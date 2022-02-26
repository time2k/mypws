package mylibs

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

//RandEngString 随机字符串生成,去掉容易歧义的和元音字母，不含数字
func RandEngString(n int) string {
	var letterBytes = []byte("bcdfghjkmnpqrstwxz")
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//RandString 随机字符串生成,去掉容易歧义的和元音字母，含数字
func RandString(n int) string {
	var letterBytes = []byte("23456789bcdfghjkmnpqrstwxz")
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//RandNumberString 随机数字字符串生成
func RandNumberString(n int) string {
	var letterBytes = []byte("0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//RandNumber 在指定范围随机输出数字
func RandNumber(ran int) int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	rd := rand.Intn(ran) //[0,n)
	return rd
}

//GoodNumber 判断给定num是否为靓号
func GoodNumber(num int) bool {
	nums := strconv.Itoa(num)
	reg := regexp.MustCompile(`(111|222|333|444|555|666|777|888|999|000){1}(111|222|333|444|555|666|777|888|999|000){1}`)
	reg2 := regexp.MustCompile(`(11111|22222|33333|44444|55555|66666|77777|88888|99999|00000){1}`)
	reg3 := regexp.MustCompile(`(11|22|33|44|55|66|77|88|99|00){1}(11|22|33|44|55|66|77|88|99|00){1}(11|22|33|44|55|66|77|88|99|00){1}`)
	if reg.MatchString(nums) {
		return true
	} else if reg2.MatchString(nums) {
		return true
	} else if reg3.MatchString(nums) {
		return true
	} else {
		return false
	}
}

//IntSize 计算整数位数
func IntSize(num int) int {
	var sizeTable = []int{
		9,
		99,
		999,
		9999,
		99999,
		999999,
		9999999,
		99999999,
		999999999,
		9999999999,
		99999999999,
		999999999999,
		9999999999999,
		99999999999999,
		999999999999999,
		9999999999999999,
		99999999999999999,
		999999999999999999,
		int(^uint(0) >> 1)}
	for i := 0; ; i++ {
		if num <= sizeTable[i] {
			return i + 1
		}
	}
}
