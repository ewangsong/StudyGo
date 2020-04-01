package main

//在utf8字符串判断是否包含指定字符串，并返回下标。 “北京天安门最美丽” , “天安门” 结果：2
import (
	"fmt"
	"strings"
)

var (
	s string = "hello北京天安门最美丽"
)

// func main() {
// 	index := strings.Index(s, "h")
// 	fmt.Println(index)
// }

// func Index(s, sub string) int {
// 	index := strings.Index(s, sub)

// 	if index == -1 || index == 0 {
// 		return index
// 	}

// }

func main() {
	fmt.Println(Utf8Index("12ws北京天安门最美丽", "天安门"))
}

func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	fmt.Println(asciiPos)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}
