# `Gorm`

### 定义结构体映射表

`GORM` 使用模型来映射一张数据库表，模型是 `struct`，由基本数据类型，实现了 `Scanner` 和 `Valuer` 接口的自定义类型及其指针组成

##### 标签

可以使用 `gorm` 字段标签来控制数据库表字段的类型，列大小，默认值等属性

### 连接数据库

支持多种数据库驱动，需要额外引入

### 自动迁移表结构

### CURD

##### 查询

- `First`
  - 默认使用主键升序排列
- `Last`
- `Where`
- `Order`
- `Limit & Offset`
  - 使用 `-1` 取消限制
- `Count`

##### 更新

- `Save`
  - 只修改结构体对象属性
- `Update`
- `Updates`
  - 只更新非零的字段

##### 删除

- `Delete`
  - 默认使用逻辑删除，实际执行的是 `Update`
  - 查询更新会自动添加 `deleted_at IS NULL`
- `Unscoped`
  - 绕过逻辑删除

##### 关联

数据库表中的关联字段的数据不会实际存储到该表中

`foreignKey` 和 `reference` 指明关联的条件，`constraint` 指明约束条件
> 生产环境不推荐使用外键，应该通过业务层面来进行约束

删除不会删除关联表的数据

##### 钩子

在创建，查询，更新，删除等操作前后调用的函数，用来关联对象的生命周期

##### 原生 SQL

`Raw` 方法执行原生的查询 `SQL`，并将结果 `Scan` 到模型中

#### 调试

- 开启全局日志
  - 创建连接时可以选择开启日志
  - `Logger:logger.Default.LogMode(logger.Info)`
- 打印慢查询 `sql`
  - `SlowThreshold: 3 * time.Millisecond`
- 打印指定 `SQL`
  - `db.Debug`
- 全局 `DryRun`
  - 连接数据库之后，语句不会真实执行
  - `DryRun: true`
- 局部 `DryRun`
  - `db.Session(&gorm.Session{DryRun: true})`