package interfaces

import "User/models"

type Iuser interface{
	Createuser(user *models.User)(string,error)
	ListRole(role *models.List)(*models.List,error)
}