- [ ] [Seata实战-分布式事务简介及demo上手](https://blog.csdn.net/hosaos/article/details/89136666)
  后面搭建未看完
- [X] [面试必问：分布式事务六种解决方案](https://zhuanlan.zhihu.com/p/183753774)
- [X] [分布式事务有这一篇就够了](https://www.cnblogs.com/dyzcs/p/13780668.html)
- [X] [漫画：什么是分布式事务？](https://blog.csdn.net/bjweimengshu/article/details/79607522)
  两阶段提交，清晰明了
- [ ] [Pattern: Event sourcing](https://microservices.io/patterns/data/event-sourcing.html)
- [ ] [Pattern: Saga](https://microservices.io/patterns/data/saga.html)
- [ ] [polling-publisher](https://microservices.io/patterns/data/polling-publisher.html)
- [ ] [Pattern: Transaction log tailing](https://microservices.io/patterns/data/transaction-log-tailing.html)


## todo
- 了解内存池设计
  - [ ] nginx ngx_pool_t
  - [ ] tcmalloc
- [ ] 缓存选型, slot, CRC16
- 缓存穿透, 下面核心思路是只让一个人来查db，或者到队列，来执行构建缓存
  - singlefly
  - 分布式锁
  - 队列
  - lease 租约
- 高级玩儿法：一些cache proxy 设计一个算法来决定一次回源的粒度



