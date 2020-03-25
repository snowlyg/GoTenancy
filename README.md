<h1 align="center">GoTenancy</h1>

<div align="center">
    <a href="https://travis-ci.org/snowlyg/go-tenancy"><img src="https://travis-ci.org/snowlyg/go-tenancy.svg?branch=master" alt="Build Status"></a>
    <a href="https://codecov.io/gh/snowlyg/go-tenancy"><img src="https://codecov.io/gh/snowlyg/go-tenancy/branch/master/graph/badge.svg" alt="Code Coverage"></a>
    <a href="https://goreportcard.com/report/github.com/snowlyg/go-tenancy"><img src="https://goreportcard.com/badge/github.com/snowlyg/go-tenancy" alt="Go Report Card"></a>
    <a href="https://godoc.org/github.com/snowlyg/go-tenancy"><img src="https://godoc.org/github.com/snowlyg/go-tenancy?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/snowlyg/go-tenancy/blob/master/LICENSE"><img src="https://img.shields.io/github/license/snowlyg/go-tenancy" alt="Licenses"></a>
    <h5 align="center">多商户管理平台</h5>
</div>

#### 项目介绍  
>
> 某次面试公司老板提到项目只是用别人的框架简单的封装，觉得没有什么技术含量（当然确实没有啥技术含量）。
> 虽然我个人认为学会使用别人设计好的框架没有什么不好，但是经过一番考虑还是决定抛弃 qor ，直接使用 iris 构建这个项目。
>
> 项目还在开发中，欢迎大家指点
---

#### 文档
> 提供一个 iris'wiki 中文文档，方便大家学习 iris : 

- [learnku.com 论坛地址](https://learnku.com/docs/iris-wiki/v12)
- [github 地址](https://github.com/snowlyg/iris/wiki)


---
也欢迎加入 Iris-go 学习交流QQ群，一起交流学习心得 ：676717248 

<a target="_blank" href="//shang.qq.com/wpa/qunwpa?idkey=cc99ccf86be594e790eacc91193789746af7df4a88e84fe949e61e5c6d63537c"><img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="Iris-go" title="Iris-go"></a>

If you don't have a QQ account, you can into the [iris-go-tenancy/community](https://gitter.im/iris-go-tenancy/community?utm_source=share-link&utm_medium=link&utm_campaign=share-link) .

[![Gitter](https://badges.gitter.im/iris-go-tenancy/community.svg)](https://gitter.im/iris-go-tenancy/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) 


#### 系统需求

- go > 1.13.0+ 
- redis 
- mysql or gcc(sqllite)


#### 配置

```shell script
  # 复制生成配置文件，并修改相关数据 
  copy /config/application.yml.example /config/application.yml

  # 修改 /config/config.go 文件的 24 行代码，将路径修改为你项目的路径
  var Root = os.Getenv("GOPATH") + "/src/github.com/snowlyg/go-tenancy"
```
 

#### 启动项目

```shell script
  # 构建数据表，重建数据。每次执行都会删除原来的数据表   
  go run ./seeder/.

  # 热启动项目 (需要安装相关工具)
  gowatch

 # 正常启动项目
 go run main.go

```
 

