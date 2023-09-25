package main

import (
	"context"
	"net/http"

	"log"

	"User/constants"
	u "User/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main(){
	r:=gin.Default()
	conn,err:=grpc.Dial(constants.Port,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("failed to connect :%v",err)
	}
	defer conn.Close()
	client:=u.NewUserserviceClient(conn)
	r.POST("/add",func(c*gin.Context){
		var request u.User
		if err :=c.ShouldBindJSON(&request);err!=nil{
			c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			return
		}
		res,err:=client.CreateUser(context.TODO(),&u.User{
			Name:     request.Name,
			Email:    request.Email,
			PhoneNo:  request.PhoneNo,
			Password: request.Password,
			Role : request.Role,
		})
		if err!= nil{
			c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		}
		c.JSON(http.StatusOK,gin.H{"message":res.Mes})
	})
	r.GET("/list",func(c*gin.Context){
		var request u.Role
		if err :=c.ShouldBindJSON(&request);err!=nil{
			c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			return
		}
		res,err:=client.ListRole(context.TODO(),&u.Role{
			Role: request.Role,
		})
		if err!= nil{
			c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		}
		c.JSON(http.StatusOK,gin.H{"role":res.Role,"Responsibility":res.Responsibility,"Access":res.Access})
	})
	// r.POST("/addrole",func(c*gin.Context){
	// 	var req u.Role
	// 	if err :=c.ShouldBindJSON(&req);err!=nil{
	// 		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	// 		return
	// 	}
	// 	res,err :=client.CreateRole(context.TODO(),&u.Role{
	// 		Name: req.Name,
	// 		Role: req.Role,
	// 	})
	// 	if err!= nil{
	// 		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	// 	}
	// 	c.JSON(http.StatusOK,gin.H{"message":res.Mes})
	// })
	r.Run(":2000")
	
}