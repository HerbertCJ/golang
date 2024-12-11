package request

type UserCreateRequest struct {
	Username string `validate:"required,min=1,max=50" json:"username"`
	Email    string `validate:"required,min=1,max=50" json:"email"`
}

type UserUpdateRequest struct {
	Id       uint   `validate:"required"`
	Username string `validate:"required,min=1,max=50" json:"username"`
	Email    string `validate:"required,min=1,max=50" json:"email"`
}