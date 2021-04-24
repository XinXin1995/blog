package request

type RegisterStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type RegisterUserRole struct {
	Id        string `json:"id"`
	IsDefault bool   `json:"isDefault"`
}

type UserBindRole struct {
	UserId string             `json:"userId" binding:"required"`
	Roles  []RegisterUserRole `json:"roles" binding:"required"`
}

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
