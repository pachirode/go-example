应用程序配置库，可以处理多种配置需求和格式

- 设置默认值
- 从多种配置文件和环境变量中读取配置信息，实时监视配置文件
- 通过命令行标志指定选项设置覆盖值

### 使用

小型项目直接使用，大型项目反序列化到结构体

##### 加载配置优先级

- 显式调用 `viper.Set`
- 命令行参数
- 环境变量
- 配置文件
- `key/value` 存储
- 默认值

##### Viper 读取配置

读取方式

- 设置默认配置值
- 从配置文件读取配置
- 监控并重新读取配置文件
- 从 `io.Reader` 读取
- 环境变量读取
- 命令行参数读取
- 远程 `key/value`
    - 需要匿名导入 `viper/remote`

```bash
docker run \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    --name=badger \
    consul:1.10.0 agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
```

##### 从 viper 中读取配置

- `Get(key string) interface{}`
    - `key` 不区分大小写
- `Get<Type>(key string) <Type>`
    - 获取指定类型配置文件
- `AllSettings() map[string]interface{}`
    - 返回所有配置
- `.`
    - 嵌套访问
- `viper.Sub()`
    - 提取子树
- 