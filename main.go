package main

import (
	"fmt"
	"github.com/zhangzw001/gin-example/pkg/setting"
	"github.com/zhangzw001/gin-example/routers"
	"net/http"
)

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
