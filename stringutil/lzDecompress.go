package stringutil

import (
	"errors"
	"math"
	"unicode/utf8"
)

var keyStrBase64Map = map[byte]int{74: 9, 78: 13, 83: 18, 61: 64, 109: 38, 114: 43, 116: 45, 101: 30, 47: 63, 73: 8, 81: 16, 113: 42, 49: 53, 50: 54, 54: 58, 76: 11, 100: 29, 107: 36, 121: 50, 77: 12, 89: 24, 105: 34, 66: 1, 69: 4, 85: 20, 48: 52, 119: 48, 117: 46, 120: 49, 52: 56, 56: 60, 110: 39, 112: 41, 70: 5, 71: 6, 79: 14, 88: 23, 97: 26, 102: 31, 103: 32, 67: 2, 118: 47, 65: 0, 68: 3, 72: 7, 108: 37, 51: 55, 57: 61, 82: 17, 90: 25, 98: 27, 115: 44, 122: 51, 53: 57, 86: 21, 106: 35, 111: 40, 55: 59, 43: 62, 75: 10, 80: 15, 84: 19, 87: 22, 99: 28, 104: 33}

type dataStruct struct {
	input      string
	val        int
	position   int
	index      int
	dictionary []string
	enlargeIn  float64
	numBits    int
}

// getBaseValue(data.input[data.index])
func getBaseValue(char byte) (int, bool) {
	i, ok := keyStrBase64Map[char]
	return i, ok
}

// Input is composed of ASCII characters, so accessing it by array has no UTF-8 pb.
func readBits(nb int, data *dataStruct, resetValue int, getNextValue func(int) int) int {
	result := 0
	power := 1
	for i := 0; i < nb; i++ {
		respB := data.val & data.position
		data.position = data.position / 2
		if data.position == 0 {
			data.position = resetValue
			data.val = getNextValue(data.index)
			data.index += 1
		}
		if respB > 0 {
			result |= power
		}
		power *= 2
	}
	return result
}

func appendValue(data *dataStruct, str string) {
	data.dictionary = append(data.dictionary, str)
	data.enlargeIn -= 1
	if data.enlargeIn == 0 {
		data.enlargeIn = math.Pow(2, float64(data.numBits))
		data.numBits += 1
	}
}

func getString(last string, data *dataStruct, resetValue int, getNextValue func(int) int) (string, bool, error) {
	c := readBits(data.numBits, data, resetValue, getNextValue)
	switch c {
	case -1:
		return "", true, errors.New("EOF")
	case 0:
		str := string(readBits(8, data, resetValue, getNextValue))
		appendValue(data, str)
		return str, false, nil
	case 1:
		str := string(readBits(16, data, resetValue, getNextValue))
		appendValue(data, str)
		return str, false, nil
	case 2:
		return "", true, nil
	}
	if c < len(data.dictionary) {
		return data.dictionary[c], false, nil
	}
	if c == len(data.dictionary) {
		return concatWithFirstRune(last, last), false, nil
	}
	return "", false, errors.New("bad character encoding")
}

// Need to handle UTF-8, so we need to use rune to concatenate
func concatWithFirstRune(str string, getFirstRune string) string {
	r, _ := utf8.DecodeRuneInString(getFirstRune)
	return str + string(r)
}

func DecompressFromBase64(input string) (string, error) {
	fromBase64, err2 := decompressFromBase64(input)
	if err2 == nil {
		return fromBase64, err2
	}
	return input, nil
}
func decompressFromBase64(input string) (string, error) {
	getNextValue := func(index int) int {
		value, b := getBaseValue(input[index])
		if !b {
			return -1
		}
		return value
	}

	value, b := getBaseValue(input[0])
	if !b {
		return input, nil
	}
	data := dataStruct{input, value, 32, 1, []string{"0", "1", "2"}, 5, 2}

	result, isEnd, err := getString("", &data, 32, getNextValue)
	if err != nil {
		return input, nil
	}
	if isEnd {
		return result, err
	}
	last := result
	data.numBits += 1
	for {
		str, isEnd, err := getString(last, &data, 32, getNextValue)
		if err != nil {
			return input, nil
		}
		if isEnd {
			return result, err
		}

		result = result + str
		appendValue(&data, concatWithFirstRune(last, str))
		last = str
	}
}
