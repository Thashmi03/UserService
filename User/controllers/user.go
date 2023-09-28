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
type Role struct{
	Roleupdate []string
}
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
func (r*RPCServer)UpdateRole(ctx context.Context,req*u.URole)(*u.UserResponse,error){
	dblist:=&models.User{
		Name:     req.Name,
		Role:     req.Role,
	}
	_,err:=UserService.UpdateRole(dblist)
	if err!=nil{
		return nil,err
	}else{
		res:= &u.UserResponse{
			Mes: "Updated",
		}
		return res,nil
	}
}
func(r*RPCServer)AddRole(ctx context.Context,req *u.ARole)(*u.UserResponse,error){
	dblist:=&models.Add{
		Name:     req.Name,
		Role:     req.Role,
	}
	_,err:=UserService.AddRole(dblist)
	if err!= nil{
		return nil,err
	}else{
		res:=&u.UserResponse{
			Mes: "role added",
		}
		return res,nil
	}

}

func (r *RPCServer) Disable(ctx context.Context, req *u.Name) (*u.Update, error) {
    dblist := &models.User{
        Name: req.Name,
    }

    users, err := UserService.Disable(dblist)
    if err != nil {
        return nil, err
    }

    // Create an instance of u.Update
    update := &u.Update{
        List: make([]*u.User, len(users)),
    }

    for i, user := range users {
        update.List[i] = &u.User{
            Name:     user.Name,
            Email:    user.Email,
            PhoneNo:  user.PhoneNo,
            Password: user.Password,
            Status:   user.Status,
            Role:     user.Role,
        }
    }

    return update, nil
}



