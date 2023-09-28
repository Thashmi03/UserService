// package main

// import (
// 	"context"
// 	// "fmt"
// 	"net/http"

// 	"log"

// 	"User/constants"
// 	u "User/proto"

// 	"github.com/gin-gonic/gin"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	r := gin.Default()
// 	conn, err := grpc.Dial(constants.Port, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("failed to connect :%v", err)
// 	}
// 	defer conn.Close()
// 	client := u.NewUserserviceClient(conn)
// 	r.POST("/add", func(c *gin.Context) {
// 		var request u.User
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
// 			return
// 		}
// 		res, err := client.CreateUser(context.TODO(), &u.User{
// 			Name:     request.Name,
// 			Email:    request.Email,
// 			PhoneNo:  request.PhoneNo,
// 			Password: request.Password,
// 			Role:     request.Role,
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": res.Mes})
// 	})
// 	r.GET("/list", func(c *gin.Context) {
// 		var request u.Role
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
// 			return
// 		}
// 		res, err := client.ListRole(context.TODO(), &u.Role{
// 			Role: request.Role,
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
// 		}
// 		c.JSON(http.StatusOK, gin.H{"role": res.Role, "Responsibility": res.Responsibility, "Access": res.Access})
// 	})
// 	r.GET("/updaterole", func(c *gin.Context) {
// 		var req u.URole
// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
// 			return
// 		}
// 		res, err := client.UpdateRole(context.TODO(), &u.URole{
// 			Name: req.Name,
// 			Role: req.Role,
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": res.Mes})
// 	})
// 	r.POST("/addrole", func(c *gin.Context) {
// 		var request u.ARole
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
// 			return
// 		}
// 		res, err := client.AddRole(context.TODO(), &u.ARole{
// 			Name: request.Name,
// 			Role: request.Role,
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": res.Mes})
// 	})
// 	r.POST("/disable", func(c *gin.Context) {
// 		var request u.Name
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		response, err := client.Disable(context.TODO(), &u.Name{
// 			Name: request.Name,
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
	
// 		// Prepare the response JSON
// 		responseData := make([]gin.H, len(response.List))
// 		for i, user := range response.List {
// 			responseData[i] = gin.H{
// 				"name":    user.Name,
// 				"email":   user.Email,
// 				"phone no": user.PhoneNo,
// 				"role":    user.Role,
// 			}
// 		}
	
// 		// Return the response JSON
// 		c.JSON(http.StatusOK, responseData)
// 	})
	
// 	r.Run(":2000")

// }
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    SetupRoutes(r)
    r.Run(":2000")
}