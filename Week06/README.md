 参考 Hystrix 实现一个滑动窗口计数器。


- 窗口包含成功、失败、超时、拒绝，但是模拟只考虑成功情况
- 窗口内超过阈值的qps就限流
- 数据结构
    + head 第一个元素
    + tail 最后一个元素
    + data 保存多少个切片
    + size 窗口长度
    + max 窗口长度内打到多少请求会限流
    + preTimestamp 第一个窗口记录的时间
    + interval 一个窗口的时间长度，默认是1000ms
    + m  一把锁
    + sum 请求窗口内的请求总量
因为没有判断队列为空的需求，所以head和tail都是指向的元素
- 包含的方法
    + GetCount 获得窗口内的请求总数
    + getCap   获得窗口的当前容量
    + IncreSuccess  增加一个成功的请求
    + getCurrent  得到当前最后一个元素
    + Add  添加请求
    + Flush 请求更新
