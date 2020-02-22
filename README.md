<h1 align="center">GoTenancy</h1>

<div align="center">
    <a href="https://travis-ci.org/snowlyg/GoTenancy"><img src="https://travis-ci.org/snowlyg/GoTenancy.svg?branch=master" alt="Build Status"></a>
    <a href="https://codecov.io/gh/snowlyg/GoTenancy"><img src="https://codecov.io/gh/snowlyg/GoTenancy/branch/master/graph/badge.svg" alt="Code Coverage"></a>
    <a href="https://goreportcard.com/report/github.com/snowlyg/GoTenancy"><img src="https://goreportcard.com/badge/github.com/snowlyg/GoTenancy" alt="Go Report Card"></a>
    <a href="https://godoc.org/github.com/snowlyg/GoTenancy"><img src="https://godoc.org/github.com/snowlyg/GoTenancy?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/snowlyg/GoTenancy/blob/master/LICENSE"><img src="https://img.shields.io/github/license/snowlyg/GoTenancy" alt="Licenses"></a>
    <h5 align="center">多商户管理平台</h5>
</div>

#### 项目介绍 
> 本来采用此项目是从 [IrisAdminApi](https://github.com/snowlyg/IrisAdminApi) 升级而来。在开发过程中遇到一个问题：gorm 的多对多关系中间表中自定义字段。
> 在搜索的过程中无意间发现了 [gorm](https://gorm.io/zh_CN/docs/index.html) 作者团队的 CMS 框架 [Qor-Admin](https://GoTenancy/libs/admin) ,文档请见 [https://doc.getqor.com/](https://doc.getqor.com/)。
> 一见到 `Qor-Admin` 框架，我就决定使用它重构我的项目了。我相信你也会爱上它的。 —— 不想偷懒的程序员无法成为一个优秀的程序员。
> `Qor-Admin` 可以单独使用，也可以和其他框架结合使用。本项目采用 `Iris + Qor-Admin` 形式开发。
> 参考项目：[qor-example](https://GoTenancy/libs/qor-example)
> 
> `Qor-Admin` 相关组件太久未更新，看来只能自己操刀了。 复制 `Qor-Admin` 相关组件到 `GoTenancy/libs` 目录下。
>
> 
---


#### iris + qor-admin 使用案例
```go
package main

import (
    "github.com/kataras/iris/v12"
    "GoTenancy/libs/admin"
)

func main() {
    app := iris.New()

    qorPrefix := "/admin"
    // create a new qor admin instance.
    q := admin.New(&admin.AdminConfig{DB: yourDB})
    // create a qor handler and convert it to an iris one with `iris.FromStd`.
    handler := iris.FromStd(q.NewServeMux(qorPrefix))
    // register routes for "/admin" and "/admin/:xxx/..."
    app.Any(qorPrefix, handler)
    app.Any(qorPrefix+"/{p:path}", handler)

    // start the server.
    app.Listen(":8080")
}

```

#### 项目目录结构
- 项目重构中....

---

#### 更新日志
[更新日志](UPDATE.MD)
---

#### 问题总结
[问题记录](ERRORS.MD)

---

#### 项目初始化

>拉取项目

```shell script

git clone https://github.com/snowlyg/GoTenancy.git

// github 克隆太慢可以用 gitee 地址:

git clone https://gitee.com/dtouyu/GoTenancy.git

```

> 加载依赖管理包 (解决国内下载依赖太慢问题)
> golang 1.13 可以直接执行：

```shell script

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

```

> 修改 `GoTenancy/config` 下的配置文件 :
> application.yml - 应用配置 
> database.yml - 数据配置 
> smtp.yml - 邮箱配置 



> 加载数据 

```shell script

go run config/db/seeds/main.go config/db/seeds/seeds.go

```

> 打包模版文件和静态文件 

```shell script

go run main.go -compile-templates=true  //模版文件


# windows 环境如果安装失败,请尝试运行：
# npm install -g node-gyp
# npm install --global --production windows-build-tools

npm install  //安装 npm 依赖
npm run build  //打包静态文件

```


> 运行项目 

[gowatch](https://gitee.com/silenceper/gowatch)
```shell script

go get github.com/silenceper/gowatch

gowatch //安装 gowatch 后才可以使用

```

当然你也可以直接使用，此方法无法热加载

```shell script

go run main.go

```

#### 账号密码
> dev@getqor.com / testing

---
##### 单元测试 
>http test

```shell script
 go test -v  //所有测试
 
 go test -run TestUserCreate -v //单个方法

// go get github.com/rakyll/gotest@latest 增加测试输出数据颜色

 gotest 
 
```

#### 演示
> http://localhost:8080
[[localhost_8080_.png]]


> http://localhost:8080/admin
[[localhost_8080_admin.png]]

---
###### Iris-go 学习交流QQ群 ：676717248
<a target="_blank" href="//shang.qq.com/wpa/qunwpa?idkey=cc99ccf86be594e790eacc91193789746af7df4a88e84fe949e61e5c6d63537c"><img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="Iris-go" title="Iris-go"></a>

