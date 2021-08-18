# 第三章学习内容
- goroutine
    - [demo1 - 内存管理相关](goroutine/g1/g1.go)
    - [demo2 - goroutine 大小](goroutine/g2/g2.go)
    - [demo3 - 线程切换](goroutine/g3/g3_test.go)
- sync 包
    - [waitGroup - 使用和注意](sync/waitGroup/s1.go) 
    - [互斥锁](sync/互斥锁/s2.go)
    - [读写锁(TODO: 书上讲解理解不到位)](sync/读写锁/s3.go)
    - Cond
        - [demo1 - 模板](sync/cond/c1/c1.go)
        - [demo2 - signal](sync/cond/c2/c2.go)
        - [demo3 - Broadcast](sync/cond/c3/c3.go)
    - Once
        - [demo1 - once.Do 相同方法](sync/once/o1/onec1.go)
        - [demo2 - once.Do 不同方法](sync/once/o2/once2.go)
        - [demo2 - once.Do 注意死锁](sync/once/o3/once3.go)
    - 池
        - [demo1 - Pool 简单示例 ](sync/pool/p1/p1.go)
        - [demo2 - Pool 池化减少内存分配 ](sync/pool/p2/p2.go)
        - [demo3 - Pool benchmark 普通连接操作耗时 ](sync/pool/p3/p3_test.go)
        - [demo3 - Pool benchmark 池化后连接操作耗时 ](sync/pool/p3/p4_test.go)
- channel
    - [基本概念和demos](channel/base.md)
- select
    - [基本概念 demos](select/select.md)
- GOMAXPROCS 控制
    - [基本概念](GOMAXPROCS/GOMAXPROCS.md)
    

