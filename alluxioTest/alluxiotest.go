package alluxioTest

import (
	"fmt"
	alluxio "github.com/Alluxio/alluxio-go"
	"github.com/Alluxio/alluxio-go/option"
	"io"
	"jasondrogba/alluxio-cacheTest/startTest"
	"log"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func ReadAlluxio(hostname string, count int, dynamic bool) {
	//TODO():将循环改为并发

	rand.Seed(int64(12345))
	for i := 1; i <= count; i++ {

		multiReadRand(i, hostname, count)
		if i == (count/2) && dynamic {
			fmt.Println("现在动态缓存策略，需要调整REPLICA到LRU")
			startTest.SwitchLRU()
			dynamic = false
		}
	}

}

func multiReadRand(i int, hostname string, count int) {
	fs := alluxio.NewClient(hostname, 39999, 0)
	index := rand.Int()
	if i <= count/2 {
		index = index % 1000
		if index > 100 {
			index = (index-100)/45 + 1
		}
	} else {
		index = index % 160
		if index > 100 {
			index = (index-100)/3 + 1
		}
	}

	pathfile := fmt.Sprintf("/%d.txt", index)
	v := 0
	//log.Println("种子：", index)
	//v := 0
	//mutex.Lock()
	exists, err := fs.Exists(pathfile, &option.Exists{})
	if err != nil {
		log.Println(err)
	}
	//log.Println(index, "文件是否存在：", exists)
	if exists {
		f, err := fs.OpenFile(pathfile, &option.OpenFile{})
		defer fs.Close(f)
		if err != nil {
			log.Println(err)
		}
		//log.Println(index, "文件打开成功")
		data, err := fs.Read(f)
		if err != nil {
			log.Println(err)
		}
		//log.Println(index, "文件读取成功")
		//mutex.Unlock()
		defer data.Close()
		content, err := io.ReadAll(data)
		if err != nil {
			log.Println(err)
		}
		//log.Println(index, "文件本地IO成功")
		v = len(content)
		log.Print(index, "文件内容长度:", v)
		//TODO():加上运行的时间，可以对比时间消耗是否有减少
	} else {
		log.Println(index, "文件不存在")
		//mutex.Unlock()
	}
}
