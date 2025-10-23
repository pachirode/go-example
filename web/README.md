# go-restful

`k8s` 整个 `REST` 路由都是基于该包，通过链式调用来构建标准的 `RESTful API` 接口

### Route

表示一条请求路由记录，内置的请求路由分发器根据 `Route` 将客户端发出的 `HTTP` 请求路由到相应的 `Handler` 处理

### WebService

路由分组，一个 `WebService` 由若干个 `Routes` 组成，并且 `WebService` 内置的路由记录都有同一个 `RootPath`
，输入输出格式，基本一致的请求数据类型等一系列的通用属性
每个 `Group` 就是提供一项服务的 `API` 集合，都会维护一个 `Version`，服务升级时，只需要通过对特定版本号更新即可

`ws.Route(ws.GET("/hello").To(hello))`

- `GET`
    - 内部调用 `RouteBuilder` 方法，构造一个 `Route`
- `To`
    - 赋值 `handler` 处理函数

### Container

表示一个 `Web Server` 服务器，由多个路由分组组成，同时还会包含若干个过滤器，一个 `http.ServeMux` 多路复用器，以及 `dispatch`
路由分发函数
不同的 `Container` 监听不同的 `ip` 地址和端口

`restful.Add(ws)`

- 将 `webService` 添加到默认 `Container`

### 过滤器

支持服务级别，路由级的请求或者响应过滤，开发者可以使用过滤器来执行常规的日志记录，验证，重定向等
提供了针对请求响应的钩子，同时还可以自定义过滤器

- 实现接口方法
    - `func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain)`
- 实现传递链
    - `chain.ProcessFilter(req, resp)`
