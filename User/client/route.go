package main

import (
	"log"

	"User/constants"
	u "User/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func SetupRoutes(r *gin.Engine) {
    conn, err := grpc.Dial(constants.Port, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer conn.Close()
    client := u.NewUserserviceClient(conn)

    r.POST("/add", AddUserHandler(client))
    r.GET("/list", ListRoleHandler(client))
    r.GET("/updaterole", UpdateRoleHandler(client))
    r.POST("/addrole", AddRoleHandler(client))
    r.POST("/disable", DisableHandler(client))
}

// Define your API handlers here (AddUserHandler, ListRoleHandler, etc.)
