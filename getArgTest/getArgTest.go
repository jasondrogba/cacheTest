package getArgTest

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func GetArgsTest() int {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("请提供一个参数")
		os.Exit(1)
	}
	// 解析参数为数字
	num, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("无效的数字参数:", args[0])
		os.Exit(1)
	}

	// 打印数字
	fmt.Println("你输入的数字是:", num)
	return num
}

func ParseArgs() (int, string, int) {
	count := flag.Int("count", 0, "The count parameter")
	policy := flag.String("policy", "", "The policy parameter")
	circle := flag.Int("circle", 0, "The circle parameter")

	flag.Parse()

	fmt.Println("Count:", *count)
	fmt.Println("Policy:", *policy)
	fmt.Println("Circle:", *circle)
	return *count, *policy, *circle
}
