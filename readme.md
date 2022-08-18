# protoc-gen-gokit-micro



## 简介

本项目是基于[go-kit-micro-service](https://github.com/leewckk/go-kit-micro-service)所提供的基础封装库，而生成的go-kit框架中所需要的协议数据结构、endpoint入口函数以及`gRPC` + `http(gin)`的transport层代码。本插件支持基于**服务端**和**客户端**代码的生成。

 ![go-kit-micro-service参考模型](https://github.com/leewckk/go-kit-micro-service/raw/master/image/go-kit-micro-service-arch.png) 

 <center>**go-kit-micro-service 参考模型**</center>



### 服务端代码生成

* `endpoint`层， 提供每项服务的入口，以及与服务之间交互所需要的业务数据结构；
* `transport`层， 提供`http` 以及`gRPC`协议的协议转换以及路由注册相关入口实现；
* `swagger`接口文档；
* 基于protobuf生成的`gRPC`底层代码；

````shell
service git:(master) $ tree
.
├── auto-gen.sh
├── configure
│   └── configure.go
├── docker
│   └── consul-zipkin
│       ├── client.json
│       ├── docker-compose.yml
│       ├── README.md
│       └── server.json
├── endpoint			## 由 protoc-gen-gokit-micro生成
│   ├── calculate
│   │   └── calculate.go
│   └── version
│       └── version.go
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── micro-service
├── pb					## 由 protoc-gen-gokit-micro生成
│   ├── calculate.proto
│   ├── golang
│   │   └── pkg
│   │       ├── calculate
│   │       │   └── calculate.pb.go
│   │       └── version
│   │           └── version.pb.go
│   ├── google
│   │   ├── api
│   │   │   ├── annotations.proto
│   │   │   └── http.proto
│   │   └── protobuf
│   │       ├── any.proto
│   │       ├── api.proto
│   │       ├── compiler
│   │       │   └── plugin.proto
│   │       ├── descriptor.proto
│   │       ├── duration.proto
│   │       ├── empty.proto
│   │       ├── field_mask.proto
│   │       ├── source_context.proto
│   │       ├── struct.proto
│   │       ├── timestamp.proto
│   │       ├── type.proto
│   │       └── wrappers.proto
│   └── version.proto
├── server
│   ├── grpc
│   │   └── server.go
│   └── http
│       └── server.go
├── service
│   ├── calculate.go
│   └── version.go
├── swagger				## 由 protoc-gen-gokit-micro生成
│   ├── calculate.swagger.json
│   └── version.swagger.json
├── transport			## 由 protoc-gen-gokit-micro生成
│   ├── gin
│   │   ├── calculate
│   │   │   ├── calculate.calculate.go
│   │   │   └── calculate.messaging.go
│   │   └── version
│   │       └── version.versionservice.go
│   └── grpc
│       ├── calculate
│       │   ├── calculate.calculate.go
│       │   └── calculate.messaging.go
│       └── version
│           └── version.versionservice.go
├── version
│   └── version.go
└── version.go

````



### 客户端代码生成

* `protocol`定义，提供与服务端交互所需要的业务层数据结构定义；
* `endpoint`，按照`gRPC`以及`HTTP`协议接入点，分别封装了各服务各函数入口，以及协议转码相关业务代码；



````shell
client git:(master) $ tree
.
├── auto-gen.sh
├── go.mod
├── go.sum
├── grpc.go
├── http.go
├── invoker			## 由 protoc-gen-gokit-micro生成
│   ├── grpc
│   │   ├── calculate
│   │   │   └── calculate.go
│   │   └── version
│   │       └── version.go
│   ├── http
│   │   ├── calculate
│   │   │   └── calculate.go
│   │   └── version
│   │       └── version.go
│   └── protocol
│       ├── calculate
│       │   └── calculate.go
│       └── version
│           └── version.go
├── main.go
├── Makefile
├── micro-client
├── pb				## 由 protoc-gen-gokit-micro生成
│   ├── calculate.proto
│   ├── golang
│   │   └── pkg
│   │       ├── calculate
│   │       │   └── calculate.pb.go
│   │       └── version
│   │           └── version.pb.go
│   ├── google
│   │   ├── api
│   │   │   ├── annotations.proto
│   │   │   └── http.proto
│   │   └── protobuf
│   │       ├── any.proto
│   │       ├── api.proto
│   │       ├── compiler
│   │       │   └── plugin.proto
│   │       ├── descriptor.proto
│   │       ├── duration.proto
│   │       ├── empty.proto
│   │       ├── field_mask.proto
│   │       ├── source_context.proto
│   │       ├── struct.proto
│   │       ├── timestamp.proto
│   │       ├── type.proto
│   │       └── wrappers.proto
│   └── version.proto
├── swagger		## 由 protoc-gen-gokit-micro生成
│   ├── calculate.swagger.json
│   └── version.swagger.json
└── version.go

````



## 快速开始





### 服务端





### 客户端









