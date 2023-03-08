package response

import "errors"

type WXBasicError struct {
	ErrCode uint   `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// 小程序登录
type JSCode2SessionResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"-"`
	UnionId    string `json:"-"`
	Token      string `json:"token"`
	WXBasicError
}

// 接口调用凭据
type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	WXBasicError
}

type PhoneInfo struct {
	PhoneNumber string `json:"phoneNumber"`
}

type UserPhoneResp struct {
	PhoneInfo `json:"phone_info"`
	WXBasicError
}

func (be WXBasicError) Error() error {
	if be.ErrCode == 0 {
		return nil
	}
	return errors.New(be.ErrMsg)
}
