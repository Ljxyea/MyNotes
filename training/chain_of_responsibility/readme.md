# 责任链
使用场景
1. 当程序需要使用不同方式处理不同种类请求，而且请求类型和顺序预先未知时， 可以使用责任链模式
该模式能将多个处理者链接成一条链路， 接收到请求后，他会询问每个处理者是否能狗对其进行处理， 这样处理者都有机会来处理请求
2. 必须按顺序执行多个处理者时， 可以使用该模式
3. 所需处理者和顺序必须在运行时进行改变， 可以使用责任链模式
4. 