package controller

import (
	"Bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
ulr     -->controller   -->    logic   -->   model
请求来了 -->控制器        -->  业务逻辑   -->   模型层的增删改查
*/


func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",nil)
}

func CreateATodo (c *gin.Context) {
	//前端页面填写待办事项，点击提交会发请求到这里
	//1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	err:=models.CreateATodo(&todo)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{ 									//3.返回响应
		c.JSON(http.StatusOK,todo )
		//c.JSON(http.StatusOK,gin.H{
		//	"code":20000,
		//	"msg":"success",
		//	"data":todo,
		//})
	}
}

func GetToDoList(c *gin.Context) {
	todoList,err:=models.GetALlTodo()
	if  err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id ,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{
			"error":"id不存在",
		})
		return
	}
	todo,err:=models.GetATodo(id)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}
	c.BindJSON(&todo)
	if err=models.UpdateATodo(todo);err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,todo)
	}

}

func DeleteATodo(c *gin.Context) {
	id,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{"error":"无效id"})
		return
	}

	if err:=models.DeleteAtodo(id);err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{id:"delete"})
	}
}