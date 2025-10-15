协程池

# gammazero/workerpool

### 源码

- `taskQueue`
  - 任务提交队列
- `workerQueue`
  - 工作协程消费队列
- `waitingQueue`
  - 等待队列

##### 创建协程池对象

使用协程创建一个任务调度器，实时任务派分，控制任务在三个队列中传递和工作协程的数量

##### 执行流程

- 任务提交，进入 `taskQueue`
- 立即转发到 
  - `waitingQueue`
  - `workerQueue`
- 任务调度器处理，决定任务最终归属
  - 队列优先，`waitingQueue` 不为空
    - 优先从 `waitingQueue` 头取出任务送给 `workeQueue`
    - 新任务进入 `waitingQueue` 尾
  - 直通
    - 直接交给工作队列
    - 数量达到协程池上限，传入 `waitingQueue`