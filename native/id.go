package native

import (
	"crypto/rand"
)

const dict = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

type idUtil struct {
	Rand64 *randIdMaker
	Rand62 *randIdMaker
	Rand32 *randIdMaker
	Rand16 *randIdMaker
	Rand10 *randIdMaker
	Rand8  *randIdMaker
}

var IdUtil = idUtil{
	Rand64: &randIdMaker{64, dict}, //0-_随机数（64位）
	Rand62: &randIdMaker{62, dict}, //0-Z随机数（62位）
	Rand32: &randIdMaker{32, dict}, //0-v随机数（32位）
	Rand16: &randIdMaker{16, dict}, //16进制随机数
	Rand10: &randIdMaker{10, dict}, //10进制随机数
	Rand8:  &randIdMaker{8, dict},  //8进制随机数
}

type randIdMaker struct {
	dictSize int
	dict     string
}

func (t *idUtil) New(dict string) *randIdMaker {
	return &randIdMaker{len(dict), dict}
}
func (t *idUtil) UUID() string {
	return t.Rand16.Rand(32)
}
func (t *idUtil) ID62(size int) string {
	return t.Rand62.Rand(size)
}
func (t *idUtil) ID32(size int) string {
	return t.Rand32.Rand(size)
}
func (t *idUtil) ID16(size int) string {
	return t.Rand16.Rand(size)
}
func (t *idUtil) NUM(size int) string {
	return t.Rand10.Rand(size)
}
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