`dyno` 主要方便操作动态对象，支持任意嵌套深度和任意嵌套组合处理这些混合类型
主要用于动态数据的解析，不需要依赖于反射，用来兼容 `encoding/json` 序列化

### Get

- 循环遍历所有路径
- 根据传入的参数类型来获取元素
  - `map[string]interface{}`
    - 对应路径元素必须为 `string`
  - `map[interface{}]interface{}`
  - `[]interface{}`
    - 必须为 `int`
  - 类型不支持

##### SGet

从嵌套的 `map[string]interface{}` 结构中，通过纯字符串（不支持切片）获取值

### Set

调用 `Get` 获取最后的值，根据值的类型进行修改

### Append

为某个对象追加一个值，需要时 `slice` 类型

### Delete

### ConvertMapI2MapS

将 `map[interface{}]interface{}` 转化为 `map[string]interface{}`