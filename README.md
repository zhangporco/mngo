# mngo

## 安装
```go
go get github.com/zhangporco/mngo
```

## 使用 

新建 App.go 文件，代码如下

```go
package main

import (
	"mngo"
	"net/http"
)

func main() {
	app := mngo.NewMngo()
	app.POST("/test/post", func(w http.ResponseWriter, r *http.Request) {
		mngo.Write(w, "POST OK")
	})
	app.GET("/test/get", func(w http.ResponseWriter, r *http.Request) {
		mngo.Write(w, "GET OK")
	})
	app.Run("8080")
}
```

## 运行
```go
go run ./App.go 
```

如果你看到如下信息，那你就启动成功了。

```go
【Mngo engine】2018/03/23 14:10:30 Mngo.go:100: Mngo starting
【Mngo engine】2018/03/23 14:10:30 Mngo.go:101: Mngo port -- 8080
```

打开浏览器，访问

```go
http://localhost:8080/test/get
```

会看到页面输出： GET OK。

至此，你可以进行任何你需要的拓展。

## 关于 Porco

Wechat: porco5555

Gmail:  zhangporco@gmail.com