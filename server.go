package main

import (
    "bytes"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "strings"
)

func startServer(buff bytes.Buffer) {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.Data(http.StatusOK, "text/html", buff.Bytes())
    })

    r.GET("/:path", func(c *gin.Context) {
        c.FileFromFS("rc"+c.Request.URL.Path, http.FS(rcFS))
    })

    server = strings.TrimPrefix(server, "http://")
    ll := strings.Split(server, ":")
    if len(ll) != 2 {
        return
    }
    ip := ll[0]
    if ll[0] == "" {
        ip = "127.0.0.1"
    }
    addr := fmt.Sprintf("http://%s:%s", ip, ll[1])
    fmt.Printf("listen and serve on %s \n", addr)
    err := r.Run(server)
    if err != nil {
        fmt.Println("listen and serve err ", err)
    }
}
