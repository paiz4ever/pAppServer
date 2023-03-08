package request

type WXLoginParam struct {
	Code string `json:"code" binding:"required"`
}

type WXBindPhoneParam struct {
	// OpenId string `json:"openid" binding:"required"`
	Code string `json:"code" binding:"required"`
}
