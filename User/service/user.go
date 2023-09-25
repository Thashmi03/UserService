package service

import (
	"User/interfaces"
	"User/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{
	UserCollection *mongo.Collection
	Rolecollection *mongo.Collection
	ctx context.Context
}

func InitService(collection *mongo.Collection,rcollection *mongo.Collection,ctx context.Context)interfaces.Iuser{
	return &UserService{collection,rcollection,ctx}
}

func(u *UserService)Createuser(user *models.User)(string,error){
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password),7)
	user.Password = string(hashedPassword)
	_,err := u.UserCollection.InsertOne(u.ctx,&user)

	if err != nil {
		return "nil", err
	}

	return "successfully user created",nil
}

func (u*UserService)ListRole(role *models.List)(*models.List,error){
	filter:=bson.M{"role": role.Role}
	var list *models.List
	res:=u.Rolecollection.FindOne(u.ctx,filter)
	err:=res.Decode(&list)
	if err!=nil{
		return nil,err
	}
	return list,nil
}
