package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	fmt.Println("init...")
	rand.Seed(time.Now().UnixNano()) // 初始化随机数的资源库，如果不执行这行，不管运行多少次都返回同样的值
}
func main() {
	fmt.Println(os.Args)
	for i := 0; i < 10; i++ {
		ret := rand.Intn(100)
		fmt.Println(i+1, ret)
	}
}
