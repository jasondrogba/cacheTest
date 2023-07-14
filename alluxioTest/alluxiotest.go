package alluxioTest

import (
	"fmt"
	alluxio "github.com/Alluxio/alluxio-go"
	"github.com/Alluxio/alluxio-go/option"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func ReadAlluxio(hostname string, count int) {
	//TODO():将循环改为并发
	//计算下面循环的时间
	startTime := time.Now()
	rand.Seed(int64(12345))
	fmt.Println("开始测试")
	for i := 1; i <= count; i++ {
		//wg.Add(1)
		//go forPrinttest(i)
		multiReadRand(i, hostname)
		//multiRead(i, hostname)
	}

	endTime := time.Now()
	fmt.Println("测试结束")
	fmt.Println("测试时间：", endTime.Sub(startTime))

}

func multiReadRand(i int, hostname string) {
	fs := alluxio.NewClient(hostname, 39999, 0)
	log.Println("第", i, "次")
	index := rand.Int() % 80
	if index > 40 {
		index = (index-40)/4 + 1
	}
	pathfile := fmt.Sprintf("/%d.txt", index)
	//v := 0
	log.Println("种子：", index)
	v := 0
	//mutex.Lock()
	exists, err := fs.Exists(pathfile, &option.Exists{})
	if err != nil {
		log.Println(err)
	}
	log.Println(index, "文件是否存在：", exists)
	if exists {
		f, err := fs.OpenFile(pathfile, &option.OpenFile{})
		defer fs.Close(f)
		if err != nil {
			log.Println(err)
		}
		log.Println(index, "文件打开成功")
		data, err := fs.Read(f)
		if err != nil {
			log.Println(err)
		}
		log.Println(index, "文件读取成功")
		//mutex.Unlock()
		defer data.Close()
		content, err := io.ReadAll(data)
		if err != nil {
			log.Println(err)
		}
		log.Println(index, "文件本地IO成功")
		v = len(content)
		log.Println(index, "文件内容长度", v)
		//TODO():加上运行的时间，可以对比时间消耗是否有减少
	} else {
		log.Println(index, "文件不存在")
		//mutex.Unlock()
	}
}

func forRead(fs *alluxio.Client) {
	for i := 1; i <= 100; i++ {
		log.Println("第", i, "次")
		index := rand.Int() % 80
		if index > 40 {
			index = (index-40)/4 + 1
		}
		pathfile := fmt.Sprintf("/%d.txt", index)
		v := 0
		fmt.Println("种子：", index)

		exists, err := fs.Exists(pathfile, &option.Exists{})
		if err != nil {
			log.Fatal(err)
		}

		if exists {
			f, err := fs.OpenFile(pathfile, &option.OpenFile{})
			defer fs.Close(f)
			if err != nil {
				log.Fatal(err)
			}
			data, err := fs.Read(f)
			if err != nil {
				log.Fatal(err)
			}
			defer data.Close()

			content, err := io.ReadAll(data)
			if err != nil {
				log.Fatal(err)
			}
			v = len(content)
			fmt.Println("长度", v)
			//TODO():加上运行的时间，可以对比时间消耗是否有减少
		}

	}
}

func multiRead(i int, hostname string) {
	fs := alluxio.NewClient(hostname, 39999, 0)
	log.Println("第", i, "次")
	index := i % 80
	pathfile := fmt.Sprintf("/%d.txt", index)
	//v := 0
	log.Println("种子：", index)
	v := 0
	//mutex.Lock()
	exists, err := fs.Exists(pathfile, &option.Exists{})
	if err != nil {
		log.Println(err)
	}
	log.Println(index, "文件是否存在：", exists)
	if exists {
		f, err := fs.OpenFile(pathfile, &option.OpenFile{})
		defer fs.Close(f)
		if err != nil {
			log.Println(err)
		}
		log.Println(index, "文件打开成功")
		data, err := fs.Read(f)
		if err != nil {
			log.Println(err)
		}
		log.Println(index, "文件读取成功")
		//mutex.Unlock()
		defer data.Close()
		content, err := io.ReadAll(data)
		if err != nil {
			log.Println(err)
		}
		log.Println(index, "文件本地IO成功")
		v = len(content)
		log.Println(index, "文件内容长度", v)
		//TODO():加上运行的时间，可以对比时间消耗是否有减少
	} else {
		log.Println(index, "文件不存在")
		//mutex.Unlock()
	}
}

func forPrinttest(i int) {
	log.Println("第", i, "次")
	defer wg.Done()
	index := rand.Int() % 80
	if index > 40 {
		index = (index-40)/4 + 1
	}
	time.Sleep(1 * time.Second)

	fmt.Println("种子：", index)

	//TODO():加上运行的时间，可以对比时间消耗是否有减少
}
