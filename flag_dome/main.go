package main

import (
	"flag"
	"fmt"
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
	admin    bool
	radiusct bool
	stop     bool
	version  bool
	c        string
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

	flag.BoolVar(&admin, "admin", false, "启动后台管理")
	flag.BoolVar(&radiusct, "radiusct", false, "启动服务")
	flag.BoolVar(&stop, "stop", false, "关闭程序")
	flag.BoolVar(&version, "version", false, "查看版本")
	flag.StringVar(&c, "c", "", "配置文件路径")
	flag.Usage = usage

	flag.Parse()

	//fmt.Println(admin, radiusct, stop, version)

	if admin {
		fmt.Println("启动后台程序")
	}
	if radiusct {
		fmt.Println("启动服务")
	}

	if stop {
		fmt.Println("关闭程序")
	}
	if version {
		fmt.Println("当前版本是0.1")
	}
	// if c == "" {
	// 	os.Exit(2)
	// } else {
	// 	fmt.Println("路径为", c)
	// }

	// a := os.Args
	// fmt.Println(a[0])

	command := exec.Command("/bin/ls")

	err := command.Run()
	pid := command.Process.Pid
	fmt.Println(err)
	fmt.Println(pid)

	// if err := command.Start(); err != nil {
	// 	fmt.Println("command failed", err)
	// }
	// err := command.Wait()
	// fmt.Println(err)

}
