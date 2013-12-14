package main

import (
	"code.google.com/p/go.exp/utf8string"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "\tThe imortant あいうえお of utf8 text\n"
	str = strings.TrimFunc(str, unicode.IsSpace)

	fmt.Printf("%s\n", str[0:len(str)/2])

	u8 := utf8string.NewString(str)
	FirstHalf := u8.Slice(0, u8.RuneCount()/2)
	fmt.Printf("%s\n", FirstHalf)
}
