package response

import (
	"blog/model"
	"time"
)

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type SysUserDetailResponse struct {
	Id        string        `json:"id"`
	Username  string        `json:"username"`
	Nickname  string        `json:"nickname"`
	Avatar    string        `json:"avatar"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAT time.Time     `json:"updatedAt"`
	Roles     []SysUserRole `json:"roles"`
}

func (u *SysUserDetailResponse) GetFomSysUser(user *model.SysUser) {
	u.Id = user.Id
	u.Username = user.Username
	u.Nickname = user.Nickname
	u.Avatar = user.Avatar
	u.CreatedAt = user.CreatedAt
	u.UpdatedAT = user.UpdatedAt
}

type SysUserRole struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefault"`
}

type LoginResponse struct {
	User      SysUserDetailResponse `json:"user"`
	Token     string                `json:"token"`
	ExpiresAt int64                 `json:"expiresAt"`
}
