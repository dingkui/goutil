package idUtil

import (
	"crypto/rand"
	"gitee.com/dk83/goutils/utils/dateUtil"
)

const dict = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

var (
	Rand64 = &randIdMaker{64, dict} //0-_随机数（64位）
	Rand62 = &randIdMaker{62, dict} //0-Z随机数（62位）
	Rand32 = &randIdMaker{32, dict} //0-v随机数（32位）
	Rand16 = &randIdMaker{16, dict} //16进制随机数
	Rand10 = &randIdMaker{10, dict} //10进制随机数
	Rand8  = &randIdMaker{8, dict}  //8进制随机数

	dateSTR = dateUtil.Layout("20060102150405000")
)

type randIdMaker struct {
	dictSize int
	dict     string
}

// New 创建随机码字典，最长有效长度64
func New(dict string) *randIdMaker {
	return &randIdMaker{len(dict), dict}
}
func UUID() string {
	return Rand16.Rand(32)
}
func ID62(size int) string {
	return Rand62.Rand(size)
}
func ID32(size int) string {
	return Rand32.Rand(size)
}
func ID16(size int) string {
	return Rand16.Rand(size)
}
func NUM(size int) string {
	return Rand10.Rand(size)
}

// RandWithTime 生成时间前缀的随机码，最短为24位
func (t *randIdMaker) RandWithTime(size int) string {
	timePrefix := dateSTR.FormatNow()
	if size < 24 {
		size = 24
	}
	return timePrefix + t.Rand(size-len(timePrefix))
}

// Rand 删除指定长度的随机码
func (t *randIdMaker) Rand(size int) string {
	result := make([]byte, size)
	dictLen := byte(t.dictSize)
	addSize := size
	move := 1
	if dictLen <= 8 {
		move = 5
	} else if dictLen <= 16 {
		move = 4
	} else if dictLen <= 16 {
		move = 4
	} else if dictLen <= 32 {
		move = 3
	} else if dictLen <= 64 {
		move = 2
	}
	moveSize := 1 << (8 - move)
	if t.dictSize < moveSize {
		addSize = size*moveSize/t.dictSize + size*t.dictSize/moveSize + 2
	}
	randValues := make([]byte, addSize)

	j := 0
	for j < len(result) {
		rand.Read(randValues)
		for _, v := range randValues {
			b := v >> move
			if b >= dictLen {
				continue
			}
			result[j] = t.dict[b]
			j++
			if j == len(result) {
				break
			}
		}
	}
	return string(result)
}
