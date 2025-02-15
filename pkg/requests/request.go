package requests

type SignUp struct {
	Username             string `json:"username" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type NewGame struct {
	RoundLength int    `json:"roundLength" validate:"number"`
	UserID      string `json:"userID"`
}
