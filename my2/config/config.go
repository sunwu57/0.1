package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type config struct {
	Path         *string
	OutfilePath  *string
	Thread       *int
	PasswordFile *string
	Passwords    []string
}

var Config config

func GetValue() {
	// 定义命令行参数
	fmt.Println("hello world")
	Config.Path = flag.String("d", "E:\\code\\go_code\\src\\test", "要解压的zip文件所在的目录")
	Config.OutfilePath = flag.String("0", "./data_output", "文件输出的路径")
	Config.Thread = flag.Int("t", 15, "解压时使用的协程数")
	Config.PasswordFile = flag.String("p", "./passwd.txt", "包含密码列表的文件路径")
	//解析命令行参数
	flag.Parse()
	// 检查参数是否为空
	if *(Config.Path) == "" {
		fmt.Println("请指定要解压的zip文件所在的目录")
		os.Exit(1)
	}
	if *(Config.PasswordFile) != "" {
		data, err := ioutil.ReadFile(*(Config.PasswordFile))
		if err != nil {
			fmt.Printf("无法读取密码文件 %s: %s\n", *(Config.PasswordFile), err)
			os.Exit(1)
		}
		Config.Passwords = strings.Split(string(data), "\n")
	} else {
		Config.Passwords = []string{"123456", "password", "123"}
	}
	fmt.Println(Config.Passwords)
	return
}
