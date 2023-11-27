package main

import (
	"backend/pkg/routes"
	"log"

	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	// 初始化全局变量
	routes.InitGlobalVariable()

	// 后台接口服务
	g.Go(func() error {
		return routes.BackendServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
