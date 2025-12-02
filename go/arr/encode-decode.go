package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder

	for _, str := range strs {
		strLen := len(str)
		b.WriteString("%")
		fmt.Fprintf(&b, "%d", strLen)
		b.WriteString("%")
		b.WriteString(str)
	}
	fmt.Println(b.String())
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	i := 0
	for i < len(encoded) {
		// %2%hi
		var strLen strings.Builder
		var r int
		for r = i + 1; i < len(encoded); r++ {
			if encoded[r] == '%' {
				r += 1
				break
			}
			strLen.WriteByte(encoded[r])
		}
		sLen, _ := strconv.Atoi(strLen.String())
		res = append(res, encoded[r:r+sLen])
		i = r + sLen
	}
	return res
}
