package models

type LoginForm struct {
	UserName string
	Password string
}

type RegisterForm struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

type ValidateToken struct {
	JWTToken string `json:"jwt_token" validate:"required"`
}
