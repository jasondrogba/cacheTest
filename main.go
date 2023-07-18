package main

import (
	"fmt"
	"jasondrogba/alluxio-cacheTest/ec2test"
	"jasondrogba/alluxio-cacheTest/getArgTest"
	"jasondrogba/alluxio-cacheTest/policyTest"
	"time"
)

func main() {

	count, _, circle := getArgTest.ParseArgs()
	instanceMap := ec2test.Getec2Instance()

	//计算下面循环的时间
	startTime := time.Now()
	fmt.Println("开始测试")

	//单次循环执行一次LRU，执行一次REPLICA，执行一个动态策略
	for i := 1; i <= circle; i++ {
		//执行完整LRU测试，600次读取
		fmt.Println("第", i, "次循环", "执行LRU策略600次")
		resultLRURemote, resultLRUUFS := policyTest.PolicyTest(instanceMap, "LRU", count)
		//执行完整REPLICA测试，600次读取
		fmt.Println("第", i, "次循环", "执行REPLICA策略600次")
		resultREPLICARemote, resultREPLICAUFS := policyTest.PolicyTest(instanceMap, "REPLICA", count)
		//执行动态策略测试，300次REPLICA策略读取，执行切换，后300次LRU
		fmt.Println("第", i, "次循环", "执行动态策略300次REPLICA，300次LRU")
		resultDynamicRemote, resultDynamicUFS := policyTest.DynamicTest(instanceMap, count)
		fmt.Println("resultLRURemote:", resultLRURemote, "resultLRUUFS:", resultLRUUFS)
		fmt.Println("resultREPLICARemote:", resultREPLICARemote, "resultREPLICAUFS:", resultREPLICAUFS)
		fmt.Println("resultDynamicRemote:", resultDynamicRemote, "resultDynamicUFS:", resultDynamicUFS)

	}
	endTime := time.Now()
	fmt.Println("测试结束")
	fmt.Println("测试时间：", endTime.Sub(startTime))

}
