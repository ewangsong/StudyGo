package main

//此程序为输入一个二进制数转换成其他进制数的转换程序
import (
	"fmt"
	"math"
	"strconv"
)

var dateNumber string

const k float64 = 2

func main() {
	fmt.Print("请输入一个二进制数：")
	fmt.Scanln(&dateNumber)
	number, isOk := checkNumber(dateNumber)
	if isOk {

		fmt.Printf("转换成整数为%d\n", number)

		fmt.Printf("转换成八进制为%o\n", number)
		fmt.Printf("转换成转换成十六进制为%x\n", number)
	} else {
		fmt.Println("输入有误,请重新输入")
	}

}

//判断长度和是否为二进制数并返回一个int
func checkNumber(s string) (number int, isOk bool) {

	dateLen := len(s)
	for i := 0; i < dateLen; i++ {
		j, err := strconv.Atoi(s[i : i+1])
		if err != nil {
			return 0, false
		}
		if j == 0 {
			continue
		} else if j == 1 {
			number += int(math.Pow(k, float64(dateLen-i)))
		} else if j != 0 && j != 1 {
			return 0, false
		}
	}
	return number, true

}
