`GMP` 调度模型中 `P` 决定同时运行的 `goroutine` 数，可以通过环境变量或者 `runtime.GOMAXPROCS(n)` 来设置

### 物理机

- `runtime.GOMAXPROCS(n)`
  - `<0`
    - 返回当前 `P` 数量，宿主机上可用的 `CPU` 核心数
  - `>0`
    - 设置为修改的值

### 容器环境

无法识别是否再容器里面，如果设置值超过实际核心数会导致频繁切换上下文，影响程序性能

# go.uber.org/automaxprocs

使用匿名的方式导入，可以自动识别容器环境并设置
