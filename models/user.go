package models

var (
	Users map[string]*User
)

type User struct {
//	ObjectId string
	Name string
	Pwd string
}

func (this User)New() User{
	return User{}
}
