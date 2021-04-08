package request

type RegisterStruct struct {
	Id       string             `json:"string"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
	Nickname string             `json:"nickname" binding:"required"`
	Avatar   string             `json:"avatar" binding:"required"`
	Roles    []RegisterUserRole `json:"roles"`
}

type RegisterUserRole struct {
	Id        string `json:"id"`
	IsDefault bool   `json:"isDefault"`
}

type UserBindRole struct {
	UserId string             `json:"userId" binding:"required"`
	Roles  []RegisterUserRole `json:"roles" binding:"required"`
}

type ValidStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
