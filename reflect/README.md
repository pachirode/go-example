# 反射

`reflect` 是 `Go` 提供的反射库，用于在运行时检查类型并操作对象

### 常用方法

- `reflect.ValueOf`
    - 可以获取值详细信息，也可以操作值
        - 获取类型
        - 修改值
- `reflect.TypeOf`
    - 获取类型的详细信息
    - 名称
    - 种类
        - 数组
        - 结构体
        - 切片

### reflect.Value

`Kind` 方法可以获取对应类型的类别
`Type` 方法可以获取 `reflect.Type`

##### 指针

如果传入 `reflect.ValueOf` 的是指针类型，需要使用 `Elem` 方法获取指针指向的值，值类型可以直接获取

使用指针可以直接使用 `Set<Type>` 方法修改字段，如果是对象会触发 `panic`

##### 获取结构体字段

- `FieldByName`
- `FieldByIndex`
- `Field`

### reflect.Type

`TypeOf` 返回信息
- 字段名字
- 包路径
- 字段类型
- 字段标签
- 字段偏移量
- 索引位置
- 匿名字段
