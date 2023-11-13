### kafka面试题

1. 为什么要使用消息中心
2. kafka的特点
3. 那种情况下需要kafka
4. kafka使用哪种消费方式, pull还是push
5. kafka和zooeeper是什么关系
6. kafka的架构
7. kafka的缓冲池满了怎么办
8. 数据传输的事务有几种
9. kafka在什么情况下会出现数据丢失
10. kafka的性能好在什么地方
11. partition的数据文件(offset, messageSize, data)
12. 数据文件分段segment (顺序读写, 分段命令, 二分查找)
13. 负载均衡 (patition会均衡分布到不同broker上面)
14. 批量发送
15. 压缩
16. 消费者设计
17. consumer group
18. 如何获取topic主题的列表
19. 生产者和消费者的命令行是什么
20. 讲讲kafka维护消费状态跟踪的方法
21. 主从同步
22. kafka判断一个节点是否活着有哪两个条件
23. kafka与传统MQ消息系统之间的关键区别
24. kafka ack的三种机制
25. 消费者如何不自动提交偏移量, 由应用提交
26. 消费者故障, 出现活锁问题如何解决
27. 如何控制消费者的位置
28. kafka分布式的情况下, 如何保证消息的顺序消费
29. kafka的高可用机制是什么
30. kafka如何减少数据丢失
31. kafka如何不消费重复数据, 比如扣款
32. 