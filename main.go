package main

import (
	"fmt"
	"net/http"
	"videOnline/conf"
	"videOnline/pkg/setting"
	"videOnline/routers"
)

func main() {
	dirs := []string{"conf"} // 设置需要释放的目录
	for _, dir := range dirs {
		// 解压dir目录到当前目录
		if err := conf.RestoreAssets("./", dir); err != nil {
			break
		}
	}
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
