# 闭包
## 匿名函数
+ 匿名函数 
  1. 简介
     + 匿名函数由一个不带函数名的函数声明和函数体组成
  2. 值类型
     + 在go中，所有的函数都是值类型.
     + 既可以作为参数传递，又可以作为返回值传递
       + 匿名函数调用有两种方法:
         + 通过返回值调用
           + ```
             func main() {
             f1, f2 := calc(2, 3)
             n1 := f1(10)
             n2 := f2()```
         + 或者进行匿名函数定义的同时调用
          + ```
             func say() {
             defer func() {
             fmt.Println("hello")
             }()
             ...
             }
``` 



  