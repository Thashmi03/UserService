package controllers

import (
	"User/interfaces"
	"User/models"
	u "User/proto"
	"context"
)

type RPCServer struct{
	u.UnimplementedUserserviceServer
}
var(
	UserService interfaces.Iuser
)
func(r *RPCServer)CreateUser(ctx context.Context,req *u.User)(*u.UserResponse,error){
	dbuser:= &models.User{
		Name:     req.Name,
		Email:    req.Email,
		PhoneNo:  req.PhoneNo,
		Password: req.Password,
		Role : req.Role,
	}
	_,err:=UserService.Createuser(dbuser)
	if err!=nil{
		return nil,err
	}else{
		res:= &u.UserResponse{
			Mes: "success",
		}
		return res,nil
	}
}

func (r*RPCServer)ListRole(ctx context.Context,req *u.Role)(*u.Roleresponse,error){
	dblist:=&models.List{
		Role: req.Role,
	}
	list,err:=UserService.ListRole(dblist)
	if err!=nil{
		return nil,err
	}else{
		res:= &u.Roleresponse{
			Role:           list.Role,
			Responsibility: list.Responsibility,
			Access:         list.Access,
		}
		return res,nil
	}
}