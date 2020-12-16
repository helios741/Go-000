按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

功能：client通过id查询用户信息，然后将用户的Age打印出来


- /cmd/server存放的是grpc服务端的代码
- /cmd/client存放是grpc客户端的代码
- /api存放的是pb文件
- /config/errcode/sql 存放的是自定义错误