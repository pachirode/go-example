命令行框架，快速生成命令行工具

### 使用

##### 添加子命令

##### 使用命令行标志

- 使用 `pflag`
- 持久标志
    - 标志可以用于其分配的命令和所有子命令
    - 全局标志一般定义在根命令
- 本地标志
    - 只使用于该指定命令
    - 父命令本地标志
        - 默认忽略父命令本地标志
        - `Command.TraverseChildren` 属性设置开启
- 必选标志

##### viper

联合 `viper` 使用

##### 参数验证

内置的 `Args` 属性提供校验功能

内置校验函数

- `NoArgs`
    - 存在命令参数报错
- `ArbitraryArgs`
    - 任意参数
- 自定义验证函数

##### Hooks

执行 `Run` 函数前后，可以执行一些钩子函数

- `PersistentPreRun`
- `PreRun`
- `Run`
- `PostRun`
- `PersistentPostRun`

##### 定义 Help

##### 未知命令建议

自带该功能，可以选择是否关闭

`Command.TraverseChildren` 需要设置为 `false`

##### shell 补全

`completion` 子命令，可以为指定的 `Shell` 生成自动补全脚本

##### 生成文档

### 手脚架

`cobra-cli` 命令行工具可以快速创建一个命令行项目