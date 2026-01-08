package id

import (
	"math"
	"strings"
)

const (
	base62Charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base          = 62 // 62 characters in the base62Charset
	zero          = "0"
	maxBaseLength = 11
	notFoundIndex = -1
)

func ToBase64(num int64) string {
	if num == 0 {
		return zero
	}

	result := make([]byte, 0, maxBaseLength)

	for num > 0 {
		remainder := num % base
		result = append(result, base62Charset[remainder])
		num = num / base
	}

	return string(reverseBytes(result))
}

func FromBase64(str string) (int64, error) {
	var result int64

	for _, char := range str { //todo: check with strings.IndexByte ?
		remainder := strings.IndexRune(base62Charset, char)
		if remainder == notFoundIndex {
			return 0, ErrInvalidBase62String
		}

		if checkOverflow(result, int64(remainder)) {
			return 0, ErrValueToLargeForBase62
		}

		result = result*base + int64(remainder)
	}

	return result, nil
}

func checkOverflow(result, remainder int64) bool {
	return result > (math.MaxInt64-remainder)/base
}

func reverseBytes(data []byte) []byte {
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}

	return data
}
