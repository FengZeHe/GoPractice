## 并发模式
- 控制程序的生命周期
- 管理可复用的资源池
- 创建可以处理任务的goroutine池

### runner
1. runner包用于展示如何使用通道来监视程序的执行时间，如果程序运行时间太长，也可以用runner包来终止程序。

### pool
1. 这个包用于展示如何使用有缓冲的通道实现资源池，来管理可以在任意数量的gorutine之间共享以及独立使用的资源。这种模式
在需要共享一组静态资源的情况（如共享数据库连接或者内存缓冲区）下非常有用。
2. 如果goroutine需要从池里得到这些资源中的一个，它可以从池里申请，使用完后归还到资源池。

### work
1. work包的目的是展示如何使用无缓冲的通道来创建一个goroutine池，这些goroutine执行并及控制一组工作，并让并发执行。这种情况下
使用无缓冲的通道要比随意指定一个缓冲区大小的有缓冲通道号，因为这种情况既不需要一个工作队列，也不需要一组goroutine配置执行。