# 问题总结 (记录的问题在 [https://github.com/qor/qor-example/issues](https://github.com/qor/qor-example/issues) 上都可以找到)

- `Error 1054: Unknown column 'basics.provider' in 'where clause' - "invalid account" error message`
    
    这个问题是登陆报错，`qor-example` 中使用了 `auth_themes` 认证主题包, 其它组件更新后这个主题包没有及时更新，导致了这个错误。
    解决方法很简单，我们只需要定义一个自己的认证主题包就可以了，实现方法可以参照 [https://github.com/snowlyg/GoTenancy/config/auth/themes](https://github.com/snowlyg/GoTenancy/config/auth/themes)。


- `Error 1170: BLOB/TEXT column 'key' used in key specification without a key length`
    
    这个错误是字段定义的问题，`github.com/qor/i18n/backends/database/database.go` 文件中 `Translation` 结构 `Key` 字段定义有问题，
    修改为 varchar(500) ，代码如下：
    
    
```go
type Translation struct {
    Locale string `sql:"size:12;"`
    Key    string `sql:"size:500;"`
    Value  string `sql:"size:4294967295"`
}
```

- [Qor-Admin](https://github.com/qor/admin)  `runtime error: invalid memory address or nil pointer dereference` 
    
    这个错误出现的原因是 [Qor-Admin](https://github.com/qor/admin) 相关组件的开发要早于 go mod 模式，所以它们都没有支持 go mod 模式下的
    前端模版文件打包。
    简单的解决方案是：将相应的组件的视图直接复制到自己项目的, `app/views/qor` 目录下。
    
    另外的方法写一个方法重新注册组件视图文件，详情可以查看 [registerview.go](utils/registerviews/registerview.go)

    
    
    
  