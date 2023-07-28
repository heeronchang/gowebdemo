package response

type Login struct {
	User  string `json:"user"`
	Name  string `json:"name"`
	Age   int32  `json:"age"`
	Token string `json:"token"`
}
