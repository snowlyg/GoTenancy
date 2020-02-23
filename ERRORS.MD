# 问题总结

1. 运行测试 go test ,报错 exec: "gcc": executable file not found in %PATH%
 系统没有安装 gcc ,或者 gcc 没有配置到系统环境变量
 安装 gcc 并设置相应系统变量

 gcc 下载地址
[http://mingw-w64.org/doku.php/download](http://mingw-w64.org/doku.php/download)


2.启动报错如下内容：
 ```shell script
   GetPermissionByNameError:record not found创建账号权限
   CreatePermissionError:Error 1366: Incorrect string value: '\xE5\x88\x9B\xE5\xBB\xBA...' for column 'display_name' at row 1GetRoleByNameErr:record not found创建角色
   
   (/home/coding/go/src/GoTenancy/backend/database/models/roles.go:108) 
   [2019-09-27 12:02:22]  Error 1366: Incorrect string value: '\xE8\xB6\x85\xE7\xBA\xA7...' for column 'display_name' at row 1 
   CreateRoleErr:Error 1366: Incorrect string value: '\xE8\xB6\x85\xE7\xBA\xA7...' for column 'display_name' at row 1
   (/home/coding/go/src/GoTenancy/backend/database/models/permission.go:99) 
   [2019-09-27 12:02:22]  Error 1366: Incorrect string value: '\xE5\x88\x9B\xE5\xBB\xBA...' for column 'display_name' at row 1 
```
 数据库字符格式不对，无法识别中文
 将数据库字符设置 utf8 字符集

3. `go run config/db/seeds/main.go config/db/seeds/seeds.go` 报错
`panic: yaml: line 8: did not find expected key`

`config/db/seeds/data` 目录下的某个 yml 文件第8行的格式有错误

