package alluxioTest

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomTest() {
	//根据冷热数据分区，生成数据
	nummaprand := make(map[int]int)
	maplengre := make(map[int]int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		index := rand.Intn(20) + 1
		if index > 10 {
			index = (index-10)/4 + 1
			nummaprand[index]++
		} else {
			nummaprand[index]++
		}
		//fmt.Print(index)
		//fmt.Print(" ")
	}
	fmt.Println(nummaprand)

	//fmt.Println()
	//随机生成数据
	// 设置随机数种子
	// 生成20个随机数字
	for i := 0; i < 100; i++ {
		// 生成1到20之间的随机数字
		num := rand.Intn(20) + 1
		maplengre[num]++
		//fmt.Print(num)
		//fmt.Print(" ")
	}
	fmt.Println(maplengre)
}

func ZipfTest() {
	rand.Seed(42)

	// 定义Zipf参数
	s := 10.1
	v := 100.0

	// 创建Zipf分布
	zipf := rand.NewZipf(rand.New(rand.NewSource(0)), s, v, uint64(v))

	// 生成100个符合Zipf分布的随机数
	for i := 0; i < 100; i++ {
		num := int(zipf.Uint64())
		if num == 0 {
			continue
		}
		fmt.Print(num)
		fmt.Print(" ")
	}
}
