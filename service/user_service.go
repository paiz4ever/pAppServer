package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pAppServer/constant"
	"pAppServer/global"
	"pAppServer/manager/wx"
	"pAppServer/models/request"
	"pAppServer/models/response"
	"pAppServer/models/table"
	"pAppServer/utils"

	u "github.com/google/uuid"
)

type CUserService struct{}

type BUserService struct{}

func (cu *CUserService) LoginS(param request.WXLoginParam) (wxResp response.JSCode2SessionResp, err error) {
	wxConfig := global.Config.WX
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	resp, err := http.Get(fmt.Sprintf(url, wxConfig.AppId, wxConfig.Secret, param.Code))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&wxResp); err != nil {
		return
	}
	if err = wxResp.Error(); err != nil {
		return
	}
	db := global.DB
	var cauth table.CUserAuth
	var uuid string
	if err := db.Where("identity_type = ? and identifier = ?", constant.WXLogin, wxResp.OpenId).First(&cauth).Error; err != nil {
		uuid = u.New().String()
		db.Create(&table.CUserAuth{
			UUID:         uuid,
			IdentityType: constant.WXLogin,
			Identifier:   wxResp.OpenId,
		})
		db.Create(&table.CUser{
			UUID:   uuid,
			Avatar: "default.jpg",
			Gender: constant.GenderUnknown,
		})
	} else {
		uuid = cauth.UUID
	}
	t, err := utils.GenerateToke(uuid)
	if err != nil {
		return wxResp, err
	}
	wxResp.Token = t
	return
}

func (cu *CUserService) EditInfoS(param request.UserEditParam, uuid string) (ncuser table.CUser, err error) {
	var cuser table.CUser
	db := global.DB
	err = db.Where("uuid = ?", uuid).First(&cuser).Error
	if err != nil {
		return
	}
	if param.Gender != cuser.Gender {
		if !param.Gender.Valid() {
			err = errors.New("illegal param")
			return
		}
		ncuser.Gender = param.Gender
	}
	if param.NickName != cuser.NickName {
		ncuser.NickName = param.NickName
	}
	err = db.Model(&cuser).Updates(&ncuser).Error
	if err != nil {
		return
	}
	return ncuser, nil
}

func (cu *CUserService) BindPhoneS(param request.WXBindPhoneParam) (wxResp response.UserPhoneResp, err error) {
	jsonBytes, err := json.Marshal(&param)
	if err != nil {
		return
	}
	token := wx.WXMgr.GetAccessToken()
	url := "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s"
	resp, err := http.Post(fmt.Sprintf(url, token), "application/json;charset=UTF-8", bytes.NewReader(jsonBytes))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&wxResp)
	return
}
