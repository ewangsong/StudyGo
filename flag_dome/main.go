package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// type MyString string

// func (m *MyString) String() string {
// 	return fmt.Sprintf("%s", *m)
// }
// func (m *MyString) Set(s string) error {
// 	*m = MyString(s)
// 	return nil
// }

var (
	start   bool
	stop    bool
	version bool
	c       string
	d       bool
)

func usage() {
	fmt.Fprintln(os.Stderr, "nginx")
	flag.PrintDefaults()
}

func main() {

	// var ip = flag.Int("flagname", 1234, "help 123")
	// var value string
	// flag.StringVar(&value, "name", "11111", "hahahah")
	// var ss MyString
	// flag.Var(&ss, "wangsong", "测试")

	// flag.Parse()

	// fmt.Println(*ip)
	// fmt.Println(value)
	// fmt.Println(ss)

	flag.BoolVar(&start, "start", false, "启动")
	flag.BoolVar(&stop, "stop", false, "关闭程序")
	flag.BoolVar(&version, "version", false, "查看版本")
	flag.StringVar(&c, "c", "", "配置文件路径")
	flag.BoolVar(&d, "d", false, "是否后台启动")
	flag.Usage = usage

	flag.Parse()

	command := exec.Command("../study")

	if start {
		err := command.Start()
		fmt.Printf("gonne start, [PID] %d running...\n", command.Process.Pid)
		ioutil.WriteFile("./gonne.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
		if err != nil {
			fmt.Println("启动程序失败", err)
			return
		}
	}

	if stop {
		fmt.Println("关闭程序")
		stopp()
	}
	if version {
		fmt.Println("当前版本是0.1")
	}

	if d {
		fmt.Println("后台启动")

	} else {
		command.Wait()

	}
	// if c == "" {
	// 	os.Exit(2)
	// } else {
	// 	fmt.Println("路径为", c)
	// }

	// a := os.Args
	// fmt.Println(a[0])
	//	stoppro()
}

func stopp() {
	b, err := ioutil.ReadFile("./gonne.lock")

	if err != nil {
		fmt.Println("获取程序PID错误", err)
		return
	}

	stopcommand := exec.Command("/bin/kill", string(b))

	stopcommand.Start()

}
