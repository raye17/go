[toc]

# io包
## 数据输入输出
### 四个核心接口
#### 1 Reader

定义了read方法，用于读取数据到字节数组中
```
	type Reader interface{
	   Read(p []byte) (n int,err,error)
	}
	数据读入到p中，返回读取的字节数n和遇到的错误error  
```
任何实现了 io.Reader 的类型都可以调用Read方法  
至于具体如何读取文件，标准库已经实现了，不用再做一遍，达到了重用的目的  
<a href="./RWer/Reader01/main.go">Reader接口示例1</a>  
<a href="./RWer/Reader02/main.go">Reader接口示例2</a>
#### 2 Writer
定义了Write方法，用于写数据到文件中
```
	type Writer interface{
	   Write(p []byte)(n int,err error)
	}
	将数据写入p中，返回成功写入的字节数n和遇到的错误error
```  
<a href="./RWer/Writer01/main.go">Writer接口示例1</a>  
<a href="./RWer/Writer02/main.go">Writer接口示例2</a>
#### 3 Closer
定义了Close方法，用于关闭连接
```
	type Closer interface {
	   Close() error
	}
```
#### 4 Seeker
定义了Seek方法，用于指定下次读取或者写入时的偏移量
```
	type Seeker interface {
	   Seek(offset int64, whence int) (int64, error)
	}
```
入参：计算新偏移量的起始值 whence， 基于whence的偏移量offset
返回值：基于 whence 和 offset 计算后新的偏移量值，以及可能产生的错误<br>
io包中定义了三种whence
```
	const (
		SeekStart =0    //基于当前文件开始位置
		SeekCurrent=1  //基于当前偏移量
		SeekEnd=2     //基于文件结束位置
```
## 文件操作
### os.File
os.File实现了io.Reader和io.Writer，因此可以在任何 io 上下文中使用。  
<a href="File/file01/main.go">向文件写入</a>  
<a href="File/file02/main.go">从文件读取</a>

### io.Copy()
io.Copy()可以轻松地将数据从一个 Reader 拷贝到另一个 Writer。  
它抽象出for循环模式并正确处理io.EOF和字节计数  
<a href="File/file03/main.go">io.Copy()函数示例</a>

### io.WriteString()
此函数让我们方便地将字符串类型写入一个 Writer  
<a href="./File/file04/main.go">示例</a>  