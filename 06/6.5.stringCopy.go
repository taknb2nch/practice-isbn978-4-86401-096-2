package main

import (
	"fmt"
)

func main() {
	str := "A string"
	bytes := make([]byte, len(str))

	// コピーされます。参照先は別
	copy(bytes, str)

	// 2回目のコピー
	strCopy := string(bytes)
	//strCopy := string([]byte(str))
	//strCopy := str

	if strCopy != str {
		panic("Copying failed!")
	}

	if &str != &strCopy {
		fmt.Println("&str != &strCopy")
	}

	fmt.Printf("%#v\n", strCopy)
}