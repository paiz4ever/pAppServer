package request

import "pAppServer/constant"

type UserEditParam struct {
	Avatar   string              `json:"-"`
	NickName string              `json:"nickname"`
	Gender   constant.UserGender `json:"gender"`
}
