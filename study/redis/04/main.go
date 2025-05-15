package main

import (
	"context"
	"fmt"
	"study/redis/cache"
	"time"
)

/*
Redis 本身没有真正的“文件夹”或目录结构，所有 key 都是平铺存储的。
但可以通过 key 的命名规范来模拟类似文件夹的层级结构，常用冒号（:）作为分隔符。例如：
- user:1001:name
- user:1001:profile
- order:2023:001
*/
func main() {
	i := 0
	client := cache.NewClientRedis()
	for {
		if i == 5 {
			break
		}
		i++
		k := fmt.Sprintf("sxy:raye00%d:test", i)
		err := client.Set(context.TODO(), k, i, time.Second*1000).Err()
		if err != nil {
			fmt.Println(err, k)
		}
		fmt.Println("set key success: ", i)
	}
	fmt.Println("set over")
}
