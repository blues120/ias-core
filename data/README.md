# Data

## ORM 框架

[ent](https://entgo.io/zh/docs/getting-started/)

## 添加模型

1. 以 User 为例，在项目根目录，执行：
    ```shell
    go run -mod=mod entgo.io/ent/cmd/ent new --target ./internal/data/schema User
    ```
2. 生成 ent 代码：`make generate`

## 单元测试

原则上尽量使用 table-driven（表格驱动测试）的方式编写单元测试。

参考文章：

- [Golang 高质量单元测试之 Table-Driven：从入门到真香](https://zhuanlan.zhihu.com/p/475314759)
- [Go Test 单元测试简明教程](https://geektutu.com/post/quick-go-test.html)
- [Go单测从零到溜系列0—单元测试基础](https://www.liwenzhou.com/posts/Go/unit-test-0/)

### DB

使用 sqlite 模拟，参考 https://github.com/ent/ent/issues/217

### Redis