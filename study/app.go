package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//	func main() {
//		redis := cache.NewClientRedis()
//		fmt.Println(redis)
//		keys, err := redis.Keys(context.TODO(), "*").Result()
//		if err != nil {
//			fmt.Println(err)
//		}
//		for _, key := range keys {
//			typeStr, err := redis.Type(context.TODO(), key).Result()
//			if err != nil {
//				fmt.Println("获取类型错误:", err)
//				continue
//			}
//			switch typeStr {
//			case "string":
//				v, err := redis.Get(context.TODO(), key).Result()
//				if err != nil {
//					fmt.Println("string类型获取失败:", err)
//				} else {
//					fmt.Println("[string] k:", key, "v:", v)
//				}
//			case "zset":
//				zvals, err := redis.ZRangeWithScores(context.TODO(), key, 0, -1).Result()
//				if err != nil {
//					fmt.Println("zset类型获取失败:", err)
//				} else {
//					fmt.Printf("[zset] k: %s v: %v\n", key, zvals)
//				}
//			case "hash":
//				hvals, err := redis.HGetAll(context.TODO(), key).Result()
//				if err != nil {
//					fmt.Println("hash类型获取失败:", err)
//				} else {
//					fmt.Printf("[hash] k: %s v: %v\n", key, hvals)
//				}
//			default:
//				fmt.Println("未处理类型:", typeStr, "key:", key)
//			}
//		}
//	}
func main() {
	// node := snowflakeNode.NewSf()
	// id := node.Generate()
	// //time.Sleep(time.Microsecond)
	// go func() {
	// 	id1 := node.Generate()
	// 	fmt.Println(id1)
	// }()
	// fmt.Println("id: ", id)
	// //fmt.Println(id1)
	// raye.Say("sss")
	// raye.Message = "HELLO"
	// raye.Say("sss")
	// sxy.Sxy()

	hashed := "$2a$10$nB45ICnl4aQCQalmIVg57.ZbmWMGlMU7Sgp5LSl/YtVWw6m3B9EKW"
	input := "admin" // 猜测的密码

	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	if err != nil {
		fmt.Println("密码不匹配")
	} else {
		fmt.Println("密码匹配 ✅")
	}

}
