package main

import (
	"Bubble/dao"
	"Bubble/models"
	"Bubble/routers"
	"log"
)




func main(){
	//创建数据库：CREATE DATABASE bubble
	//连接数据库
	err:=dao.InitMySQL()
	if err != nil {
		log.Println("connect mysql failed,err:",err)
		return
	}
	defer dao.Close()   //程序推出关闭数据库连接

	//模型与数据库中的表对应起来
	dao.DB.AutoMigrate(&models.Todo{})

	r:=routers.SetupRouter()
	err=r.Run()
	if err != nil {
		log.Println("http server connect failed,err:",err)
		return
	}

}
