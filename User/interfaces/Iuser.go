package interfaces

import "User/models"

type Iuser interface{
	Createuser(user *models.User)(string,error)
	ListRole(role *models.List)(*models.List,error)
	UpdateRole(role *models.User)(string,error)
	AddRole(role *models.Add)(string,error)
	Disable(role *models.User)([]*models.User,error)
}