package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
构建一个猜数字游戏。在这个游戏里面，程序首先会生成一个介于 1到100之间的随机整数，然后提示玩家进行猜测。
玩家每次输入一个数字，程序会告诉玩家这个猜测的值是高于还是低于那个秘密的值。如果猜对了，就告诉玩家胜利并且退出程序。
*/
func main() {
	secretNumber := rand.Intn(100)
	fmt.Println("please input your guess:")
	for {
		var guess int
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("invalid input,please input again: ", err)
			continue
		}
		fmt.Println("your guess is: ", guess)
		if guess > secretNumber {
			fmt.Println("your guess is bigger than the secret number.please try again")
		} else if guess < secretNumber {
			fmt.Println("your guess is litter than th secret number.please try again")
		} else {
			fmt.Println("good,you legend!")
			break
		}
	}
	fmt.Println("the secret number is: ", secretNumber)
}
func init() {
	fmt.Println("Initializing...")
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Initialization completed!")
}
