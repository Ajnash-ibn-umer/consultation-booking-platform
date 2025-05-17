package dto

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}
