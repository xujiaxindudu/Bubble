package routers

import (
	"Bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter()*gin.Engine{
	//创建默认路由
	r:=gin.Default()
	//告诉gin去哪里找静态文件
	r.Static("/static","static")
	//告诉gin去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	//读取前端页面
	r.GET("/", controller.IndexHandler)

	//创建一个路由组v1
	v1Group:=r.Group("v1")
	{
		//待办事项
		//添加事项
		v1Group.POST("/todo",controller.CreateATodo)
		//查看所有待办事项
		v1Group.GET("/todo", controller.GetToDoList)
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id",controller.DeleteATodo)
	}
	return r
}
