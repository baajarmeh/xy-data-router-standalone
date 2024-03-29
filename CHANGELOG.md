# XY.DataRouter

## v1.100.55.21102323

- 优化对象复用

## v1.100.54.21102111

- 减少不必要的指针间接引用
- 全局时间优化

## v1.100.50.21091809

- 调整 `main` 包结构, 增加版本信息
- 优化 `Makefile` 
- 优化 `init()` 减少内部相互依赖, 明确初始化顺序

## v1.100.45.21090909

- 内嵌 `favicon.ico`
- 增加 `/client_ip` `/server_ip` 显示客户端来访 IP 和当前服务器 IP
- 升级依赖包

## v1.100.44.21082808

- 使用 `xsync.NewMap` 替代 `cmap.New`, `xsync.Counter` 替代 `atomic`
- 接口不存在日志记录客户端 IP

## v1.100.43.21082020

- 升级 `go1.17` `utils 0.2.0` 等

## v1.100.42.21080808

- 稳定版本
- 升级依赖包, 规范注释

## v1.100.41.21071515

- 增加 `TunnelDataTotal` `TunnelTodoSendCount` 等, 规范统计字段名称
- 增加配置项 `TunSendQueueSize` 发送数据最大队列长度, 默认: `8192`, 公网大丢包可能超过队列长度而丢弃数据
- `Tunnel` 数据发送不设置超时

## v1.100.40.21071323

- 数据代理由 `WsHub(Websocket)` 替换为 `Tunnel(aRPC)`

## v1.100.31.21070417

- 新增: 无限缓冲信道最大缓冲数量配置(可选) `data_chan_max_buf_cap`
  - 默认为 `0`, 无限大
  - 当前配置 `500000`, 超过限制 `DataChanSize(初始化大小, 默认 50) + DataChanMaxBufCap` 丢弃数据
  - 消费长时间慢于生产, 避免累积数据占用内存过大, 不应该出现该状态 (`WsHubQueueDiscards`)

## v1.100.30.21070323

- 新增: 多级数据代理功能, 解决海外数据上报网络问题
  1. 程序相同, 部署到各地, 可选开启代理 `-f 1.2.3.4:1234`
  2. 就近响应 HTTP/UDP 数据请求 (需配合智能 DNS, 域名解析到就近服务器)
  3. 走专线/环网将数据传递给上联服务器 `1.2.3.4` (可多级传递)
  4. 上联服务器接收到数据后, 在本地继续完成剩下的数据处理和分发逻辑

## v1.100.18.21063012

- 修正: ES 搜索接口中指针在 `go-json` 引发异常

## v1.100.17.21062911

- 增加: 数据处理待办计数等统计项
- 优化: 数据处理创建工作单元时, 排队情况下也不阻塞

## v1.100.16.21062717

- 增加可选配置: `reduce_memory_usage` 减少内存占用(可能增加CPU占用), 默认关闭

## v1.100.15.21062616

- ES 批量写入日志级别配置方式调整
- 增加接口请求黑名单功能

## v1.100.14.21062300

- 优化组路由代码, 严格路由模式, 默认回应请求后立即关闭连接
- 升级依赖到最新, 扩充时间轮精度
- 增加 HTTP 错误请求计数

## v1.100.10.21060500

- 批量提交支持更多数据形式, `/v1/apiname/bulk` 可 POST 多条 JSON 数据 

  1. 现支持数据间用 `=-:-=` 分隔, 如: `{"a":1}=-:-={"b":2}`

  2. 增加支持列表数据, 如: `[{"a":1},{"b":2}]`

  3. 增加支持混搭,如: `{"a":1}=-:-={"b":2}=-:-=[{"c":3},{"d":4}]`

  4. 任意 JSON 格式, 有无美化样式, 换行, 空格都支持, 如:

     ```json
      {
         "a" :
       1 , "b":[
           2,
           3
       ]
     } =-
     :- =[ {"c" :"d" }, {"e":5}] =-:-=
     ```

- 优化 HTTP 接口成功返回值

- 调整初始化顺序, 守护进程更干净

- 启用时间轮, 优化定时器

- 调整关闭写入 ES 时代码逻辑, 数据处理允许的并发最小值调整: `10`

## v1.100.8.21052315

- init
