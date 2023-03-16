package v1

import (
	"fmt"
	"pAppServer/models/request"
	"pAppServer/models/response"
	"pAppServer/service"

	"github.com/gin-gonic/gin"
)

type CUserApi struct {
	service.CUserService
}

type BUserApi struct{}

func (cu *CUserApi) Login(c *gin.Context) {
	var param request.WXLoginParam
	if err := request.Vertify(&param, c); err != nil {
		return
	}
	if wxResp, err := cu.LoginS(param); err != nil {
		response.Failed("login fail", c)
		return
	} else {
		response.Success(wxResp, c)
	}
}

func (cu *CUserApi) EditInfo(c *gin.Context) {
	var param request.UserEditParam
	if err := request.Vertify(&param, c); err != nil {
		return
	}
	uuid := c.GetString("uuid")
	if uuid == "" {
		response.Failed("uuid empty", c)
		return
	}
	if cuse, err := cu.EditInfoS(param, uuid); err != nil {
		response.Failed(err.Error(), c)
		return
	} else {
		response.Success(cuse, c)
	}
}

func (cu *CUserApi) Upload(c *gin.Context) {
	r, _ := c.MultipartForm()
	for k, f := range r.File["file"] {
		fmt.Println("k===", k, "f===", f.Filename)
	}
	fmt.Println("v===", r.Value)

	fmt.Println("mmm===", c.PostForm("mmm"))
	fmt.Println("aaa===", c.PostForm("aaa"))
	fmt.Println("bbb===", c.PostForm("bbb"))
	response.Success(c.PostForm("aaa"), c)
	// file, err := c.FormFile("file")
	// if err != nil {
	// 	response.Failed("图片上传出错", c)
	// }
	// err = c.SaveUploadedFile(file, "./assets/avatar/"+file.Filename)
	// if err != nil {
	// 	fmt.Println("err", err)
	// 	return
	// }
	// response.Success(file.Filename, c)
}

func (cu *CUserApi) BindPhone(c *gin.Context) {
	var param request.WXBindPhoneParam
	if err := request.Vertify(&param, c); err != nil {
		return
	}
	if wxResp, err := cu.BindPhoneS(param); err != nil {
		response.Failed("bindphone fail", c)
		return
	} else {
		response.Success(wxResp, c)
	}
}
