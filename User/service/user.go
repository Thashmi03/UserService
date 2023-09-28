package service

import (
	"User/interfaces"
	"User/models"
	"context"
	"log"

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
	user.Status=1
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
func (u*UserService)UpdateRole(role *models.User)(string,error){
	filter := bson.D{{Key: "name", Value: role.Name}}
	fv :=bson.D{{Key: "$set",Value: bson.D{{Key: "role",Value: role.Role}}}}

	_, err := u.UserCollection.UpdateOne(context.TODO(), filter, fv)
	if err!=nil{
		return "nil",err
	}
	return "updated",nil
}
func ( u *UserService)AddRole(role *models.Add)(string,error){
	filter := bson.M{
        "name": role.Name,
    }
	update:=bson.M{
		"$push":bson.M{
			"role":role.Role,
		},
	}
	_,err:=u.UserCollection.UpdateOne(u.ctx,filter,update)
	if err!= nil{
		return "nil",err
	}
	return "added",nil
}
func (u *UserService)Disable(data *models.User)([]*models.User,error){
	filter:=bson.M{
		"name":data.Name,
	}
	update:=bson.M{
		"$set":bson.M{
			"status":0,
		},
	}
	_, err := u.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
        log.Fatal(err)
    }
	var find []*models.User
	fil := bson.M{
        "status": bson.M{
            "$ne": 0,
        },
    }
	res,_:=u.UserCollection.Find(u.ctx,fil)
	for res.Next(u.ctx){
		mod :=&models.User{}
		e:=res.Decode(mod)
		if e!=nil{
			return nil,e
		}
		find=append(find, mod)
	}

	return find,nil

}
