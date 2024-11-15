# Biz

## 接口命名规范

| 类别   | 功能           | repo层（data层） | biz层        |
|------|--------------|--------------|-------------|
| 创建资源 | 创建资源         | Save         | Create      |
| 删除资源 | 根据id删除       | Delete       | Delete      |
|      | 根据id批量删除     | BatchDelete  | BatchDelete |
| 更新资源 | 根据id更新       | Update       | Update      |
| 查询资源 | 根据id查询单条记录   | Find         | Find        |
|      | 根据A字段查询单条记录  | FindByA      | FindByA     |
|      | 根据过滤条件查询多条记录 | List         | List        |

