# 分布式锁测试

使用go-micro下的sync组件，使用基于etcd实现的sync组件接口插件

## 运行

一个窗口运行`go run main.go -o`，该程序会每一秒输出一个数字

另外一个窗口运行`go run main.go`，该程序会等待输入任意的数字，输入第一次，会上锁，第二次会解锁，由此观察第一个窗口的输出