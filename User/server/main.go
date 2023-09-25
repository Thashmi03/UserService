package main

import (
	"User/config"
	"User/constants"
	"User/controllers"
	"User/service"
	"context"
	"fmt"
	"net"
	u "User/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Initdatabase(client * mongo.Client){
	UserCollection :=config.GetCollection(client, "User", "db")
	Rolecollection :=config.GetCollection(client,"User","Role")
	controllers.UserService =service.InitService(UserCollection,Rolecollection,context.Background())
}

func main(){
	mongoclient,err:= config.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil{
		panic(err)
	}
	Initdatabase(mongoclient)
	lis,err:=net.Listen("tcp",constants.Port)
	if err!=nil{
		fmt.Println("failed :",err)
		return
	}
	server :=grpc.NewServer()
	u.RegisterUserserviceServer(server,&controllers.RPCServer{})
	fmt.Println("server listening on : ",constants.Port)
	if err := server.Serve(lis);err!=nil{
		fmt.Println("failed")
	}
}