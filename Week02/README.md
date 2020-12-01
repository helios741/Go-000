我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


答：
- 在dao层会进行wrap，因为和底层交互的要进行wrap
- 抽象了一个公用的SqlIsNotFount的方法(utils/sql.go)，为了以后切换数据库引擎，返回的Notfound的错误不同了，对上层造成大面积影响
- 上层调用的时候，使用公用的SqlIsNotFount方法
- 中间传递错误的时候，可以使用withMessage，如果不想保存信息，只保存调用栈，则使用withStack
