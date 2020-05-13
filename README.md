# gRPC+gRPC-gateway Demo App in Golang

## gRPC简介

- 基于HTTP/2传输协议的RPC框架
- 基于Protobuf定义结构化的数据模型并生成代码
- 支持插件
- 跨平台、跨语言、高性能

## 业务背景

- 单个高可用集群
- 后端业务复杂，大量微服务模块、通讯量大
- 前端页面为后台的终端

## 方案

- 提供高吞吐量的通讯服务及结构化的服务定义
- 提供一套平台标准化、结构化、可复用的接口定义
- gRPC + gRPC-gateway

## TODO

- 使用ConfigMap启动服务
- 接入jaeger tracing中间件到gRPC server、http server
