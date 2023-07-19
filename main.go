package main

import (
	"fmt"
	"jasondrogba/alluxio-cacheTest/ec2test"
	"jasondrogba/alluxio-cacheTest/getArgTest"
	"jasondrogba/alluxio-cacheTest/policyTest"
	"time"
)

const N = 10

func main() {

	_, _, circle := getArgTest.ParseArgs()
	instanceMap := ec2test.Getec2Instance()

	//计算下面循环的时间
	startTime := time.Now()
	fmt.Println("开始测试")

	//单次循环执行一次LRU，执行一次REPLICA，执行一个动态策略
	for count := 400; count <= N*100; count += 100 {
		for i := 1; i <= circle; i++ {
			//执行完整LRU测试，600次读取
			fmt.Println("第", i, "次循环", "执行LRU策略", count, "次")
			resultLRURemote, resultLRUUFS := policyTest.PolicyTest(instanceMap, "LRU", count)
			//执行完整REPLICA测试，600次读取
			fmt.Println("第", i, "次循环", "执行REPLICA策略", count, "次")
			resultREPLICARemote, resultREPLICAUFS := policyTest.PolicyTest(instanceMap, "REPLICA", count)
			//执行动态策略测试，300次REPLICA策略读取，执行切换，后300次LRU
			if count >= 600 {
				fmt.Println("第", i, "次循环", "执行动态策略", count/2, "次REPLICA，", count/2, "次LRU")
				resultDynamicRemote, resultDynamicUFS := policyTest.DynamicTest(instanceMap, count)
				fmt.Println(count, "次文件读取", "resultDynamicRemote:", resultDynamicRemote, "resultDynamicUFS:", resultDynamicUFS)
			}
			fmt.Println(count, "次文件读取", "resultLRURemote:", resultLRURemote, "resultLRUUFS:", resultLRUUFS)
			fmt.Println(count, "次文件读取", "resultREPLICARemote:", resultREPLICARemote, "resultREPLICAUFS:", resultREPLICAUFS)
		}
	}
	endTime := time.Now()
	fmt.Println("测试结束")
	fmt.Println("测试时间：", endTime.Sub(startTime))

}
