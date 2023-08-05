package main

import (
	"Douyin/cache"
	"Douyin/config"
	"Douyin/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
)

func main() {
	var err error
	err = config.ConfInit() //初始化配置文件
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = cache.RedisPoolInit() //初始化redis
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = dao.DbInit() //初始化mysql
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	r := gin.Default()   //初始化gin
	RouterInit(r)        //初始化路由
	err = r.Run(":8000") //启动并监听端口
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
