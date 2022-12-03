# Gin 中的Binding
在model中可以查看到主要有几个参数
**required**
表示这个字段必须传入
在`Email` 这个字段还发现`required`后面有个`email`，表示该字段进行邮箱校验。

**eqfield=FieldName**
表示当前字段和FieldName必须相对

# 数据库
## 连接数据库

```go
package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

func Init() (err error) {
	// Open("database type", "username:password@tcp(host:port)/dbName")
	DBConn, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_stu")
	return err
}
```
> 注意：
> 1. 在连接数据库时，记得添加数据库驱动的包
> 2. 注意`DBConn, err := sql.Open()` 是会出现空的数据库连接对象的。

## 参数介绍
sql.DB 对象是许多数据库连接的池，其中包含 ' 使用中 ' 和' 空闲 ' 两种连接状态。当您使用连接来执行数据库任务 (例如执行 SQL 语句或查询行) 时，该连接会被标记为正在使用中。任务完成后，连接将被标记为空闲。
当您指示 sql.DB 执行数据库任务时，它将首先检查池中是否有可用的空闲连接。如果有一个可用，那么 Go 将重用此现有连接并将其标记为在任务执行期间处于使用状态。如果需要时池中没有空闲连接，则 Go 将创建一个新的附加连接。

### SetMaxOpenConns
默认情况下，同时打开的连接数 (使用中 + 空闲) 没有限制。但是您可以通过 SetMaxOpenConns 方法实现对连接数的限制
```go
//设置同时打开的连接数(使用中+空闲)
//设为5。将此值设置为小于或等于0表示没有限制
//最大限制(这也是默认设置)。
db.SetMaxOpenConns(5)
```

### SetMaxIdleConns
默认情况下 sql.DB 会在链接池中最多保留 2 个空闲链接。可以通过 SetMaxIdleConns() 方法更改此方法
```go
// 将最大并发空闲链接数设置为 5.
// 小于或等于 0 表示不保留任何空闲链接.
db.SetMaxIdleConns(5)
```

# viper
## 使用
```go
viper.SetConfigFile("filePath")
viper.ReadInConfig
viper.Unmarshal(&Conf)
```
> **注意**
> 1. 配置文件读取的路径，一定是以 main.go 开始的相对路径
> 2. 注意struct 写书