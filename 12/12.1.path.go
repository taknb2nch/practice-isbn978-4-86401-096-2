package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	components := []string{"a", "path", "..", "with", "relative", "elements"}

	// スライスを可変長引数に渡します。
	path := path.Join(components...)

	fmt.Printf("Path: %s\n", path)

	decomposed := filepath.SplitList(path)

	for _, dir := range decomposed {
		fmt.Printf("%s%c", dir, filepath.Separator)
	}

	fmt.Printf("\n")
}
