package main

import (

    u "User/proto"
    "github.com/gin-gonic/gin"
)

func AddUserHandler(client u.UserserviceClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Your add user handler code
    }
}

func ListRoleHandler(client u.UserserviceClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Your list role handler code
    }
}

func UpdateRoleHandler(client u.UserserviceClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Your update role handler codegir 
    }
}

func AddRoleHandler(client u.UserserviceClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Your add role handler code
    }
}

func DisableHandler(client u.UserserviceClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Your disable handler code
    }
}
