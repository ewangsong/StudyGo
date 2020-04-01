package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	fmt.Printf("%p", &x)
	defer func(x int) {
		fmt.Printf("%p\t", &x)
		x++
	}(x)
	return 5
}
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

/*考点:defer和函数返回值 需要明确一点是defer需要在函数结束前执行。函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数
DeferFunc1有函数返回值t作用域为整个函数，在return之前defer会被执行，所以t会被修改，返回4; DeferFunc2函数中t的作用域为函数，返回1; DeferFunc3返回3
*/
/*
5
6
5
0xc00008a0180xc000090008        5
4
1
3
*/
