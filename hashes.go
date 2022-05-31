package customhashes

import (
	"strconv"
)

type HashFuncType func(input string) uint64

func Hash1(input string) uint64 {
	o1 := ""
	for _, char := range input {
		o1 += strconv.Itoa(int(char))
	}

	var o2 string
	for i := 0; i < len(o1); i++ {
		o2 += strconv.Itoa(int(o1[i]) * int(o1[len(o1)-i-1]))
	}

	o3, _ := strconv.ParseUint(o2, 10, 64)
	for _, char := range o2 {
		o3 += uint64(char)
		o3 *= uint64(char)/20 + 1
	}

	return o3
}
