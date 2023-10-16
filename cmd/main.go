package main

import (
	"github.com/CocaineCong/gin-mall/conf"
	"github.com/CocaineCong/gin-mall/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
