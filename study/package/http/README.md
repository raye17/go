# http包
## client.go源码
+ 定义了客户端对象，以及客户端send发送请求获取响应方法（调用了内部send）方法，获取截止时间，获取往返处理器方法
```
	type Client struct {
	// Transport是一个接口，表示HTTP事务，用于处理客户端的请求并等待服务端的响应。
	// 如果Transport为nil，则使用DefaultTransport。
	Transport RoundTripper
	// CheckRedirect指定处理重定向的策略（函数类型）
	// 如果CheckRedirect不为nil，客户端会在执行重定向之前调用本函数字段。
	// 参数req和via是将要执行的请求和已经执行的请求（切片，越新的请求越靠后）。
	// 如果CheckRedirect返回一个错误，本类型的Get方法不会发送请求req，
	// 而是返回之前得到的最后一个回复和该错误。（包装进url.Error类型里）
	// 如果CheckRedirect为nil，会采用默认策略：连续10此请求后停止。
	// 函数指定处理重定向的策略。当使用 HTTP Client 的Get()或者是Head()方法发送 HTTP 请求时，若响应返回的状态码为 30x （比如 301 / 302 / 303 / 307）， HTTP Client会在遵循跳转规则之前先调用这个CheckRedirect函数
	CheckRedirect func(req *Request, via []*Request) error

	// Jar指定cookie管理器：将相关cookie插入到每个出站请求中，并用每个入站响应的cookie值进行更新
	// 如果Jar为nil，请求中不会发送cookie，回复中的cookie会被忽略。（只有在请求中显式设置了cookie时，才会发送cookie）
	Jar CookieJar

	// Timeout指定本类型的值执行请求的时间限制。
	// 该超时限制包括连接时间、重定向和读取回复主体的时间。
	// 计时器会在Head、Get、Post或Do方法返回后继续运作并在超时后中断回复主体的读取。
	// Timeout为零值表示不设置超时。
	Timeout time.Duration
}
```
