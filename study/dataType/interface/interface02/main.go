package main

import (
	"fmt"
	"os"
)

// 求各种图形面积
type rectangle struct { //定义矩形
	long   int
	width  int
	height int
}
type square struct { //定义正方形
	side int
}
type triangle struct {
	height float64
	bottom float64
} //三角形
//	type circle struct {
//		radius int
//	}
type TuXing interface {
	areaV()
}

func (r rectangle) areaV() {
	var l = r.long
	var w = r.width
	var h = r.height
	var S = l * w
	var V = S * h
	fmt.Printf("矩形面积为：%5d     矩形体积为：%d\n", S, V)
}
func (s square) areaV() {
	var si = s.side
	var S = si * si
	var V = S * si
	fmt.Printf("正方形面积为：%5d   正方体体积为：%d\n", S, V)
}
func (t triangle) areaV() {
	var h = t.height
	var b = t.bottom
	var S = h * b
	fmt.Printf("三角形面积为：%.4f\n", S*1/2)
}

/*
	func (c circle) area_v() {
		var r = c.radius
		var S   = 3.14 * r * r
		var V = 4/3 * 3.14 * r * r * r
		fmt.Printf("圆面积为：%f    圆体积为：%f", S, V)
	}
*/
func count(x TuXing) {
	x.areaV()
}
func main() {
	// var r1 rectangle
	// r1.long = 4
	// r1.width = 5
	// r1.height = 6
	// var s1 = square{
	// 	side: 5,
	// }
	// var t1 = triangle{4, 6}
	// count(r1)
	// count(s1)
	// count(t1)
	for {
		fmt.Println("请选择你要计算的图形类型")
		fmt.Println(`
	1,矩形
	2,正方形
	3,三角形
	4,退出
	`)
		var choice, a, b, c int
		var d, e float64
		fmt.Scanln(&choice)
		fmt.Println("你选择的是：", choice)
		switch choice {
		case 1:
			fmt.Println("请输入长宽高")
			fmt.Scanln(&a, &b, &c)
			var r1 = rectangle{
				long:   a,
				width:  b,
				height: c,
			}
			count(r1)
		case 2:
			fmt.Println("请输入正方形边长：")
			fmt.Scanln(&a)
			var s1 = square{
				side: a,
			}
			count(s1)
		case 3:
			fmt.Println("请输入三角形底和高：")
			fmt.Scanln(&d, &e)
			var t1 = triangle{
				height: d,
				bottom: e,
			}
			count(t1)

		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入无效,请重新输入：")

		}
	}
}
