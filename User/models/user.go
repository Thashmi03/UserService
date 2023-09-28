package models

type User struct {
	Name     string   `json:"name,omitempty" bson:"name"`
	Email    string   `json:"email,omitempty" bson:"email"`
	PhoneNo  int64    `json:"phone_no,omitempty" bson:"phone_no"`
	Password string   `json:"password,omitempty" bson:"password"`
	Status   int32    `json:"status,omitempty" bson:"status"`
	Role     []string `json:"role,omitempty" bson:"role"`
}
type Add struct {
	Name     string `json:"name,omitempty" bson:"name"`
	Role     string `json:"role,omitempty" bson:"role"`
}
type List struct {
	Role           string `json:"role,omitempty" bson:"role"`
	Responsibility string `json:"responsibility,omitempty" bson:"res"`
	Access         string `json:"access,omitempty" bson:"access"`
}
