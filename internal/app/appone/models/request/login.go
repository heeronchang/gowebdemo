package request

type Login struct {
	User string `json:"user" form:"user" binding:"required,min=5,max=20"`
	Pswd string `json:"pswd" form:"pswd" binding:"required,min=5,max=20"`
}
