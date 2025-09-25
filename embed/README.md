# Embed

### go:embed

编译指令，能给在程序编译时在二进制文件中嵌入任意文件或目录，只能用在包级别的全局变量中

使用场景
- 托管静态资源服务器
- 单元测试文件依赖

##### 嵌入方式

- 将文件嵌入到 `string`
  - 适合单个文件，文本数据
  - 限制为可读
- 将文件嵌入到 `[]byte`
  - 二进制文件，图片等非文本的数据
  - 可以修改
- 嵌入到 `embed.FS`
  - 一个只读的文件系统，可以嵌入整个文件夹

##### patterns 规则

- 不支持嵌入空目录，也不支持嵌入符号链接
- 不能匹配特殊符号 ```" * < > ? ` ' | / \```
- 以 `.` 或者 `_` 开头的文件会被忽略
  - `//go:embed all:testdata` 添加 `all` 可以加入
  - `*` 也可以，但是其不具备递归性

##### `embed.FS`

- `ReadFile`
  - 读取文件中的内容
- `fs.ReadDir`
  - 读取指定文件夹下的所有文件信息
- `io/fs.FS`
  - 实现接口，可以直接转换 `http.FS`

### go list

可以查看嵌入的文件信息
```bash
# 所有被直接嵌入的 patterns，包括目录和文件
$ go list -f '{{.EmbedPatterns}}'
[testdata]
# 嵌入的文件
$ go list -f '{{.EmbedFiles}}'   
[testdata/test.txt]
# 测试文件中嵌入的 patterns
$ go list -f '{{.TestEmbedPatterns}}' 
[testdata/test.txt]
# 黑盒测试文件中嵌入的 patterns
$ go list -f '{{.XTestEmbedPatterns}}'
[testdata/test.txt]
```

