package wx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pAppServer/global"
	"pAppServer/models/response"
	"sync"
	"time"
)

type WXManager struct {
	mu          sync.RWMutex
	accessToken string
}

func NewWXManager() *WXManager {
	return new(WXManager)
}

var WXMgr = NewWXManager()

func (w *WXManager) GetAccessToken() string {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.accessToken
}

func (w *WXManager) ReqAccessToken() {
	w.mu.Lock()
	defer w.mu.Unlock()
	wxConfig := global.Config.WX
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	resp, err := http.Get(fmt.Sprintf(url, wxConfig.AppId, wxConfig.Secret))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var wxResp response.AccessTokenResp
	json.NewDecoder(resp.Body).Decode(&wxResp)
	if wxResp.ErrCode != 0 {
		fmt.Println("获取token错误", wxResp.ErrMsg)
		return
	}
	// fmt.Println("Token: ", wxResp.AccessToken, wxResp.ExpiresIn)
	w.accessToken = wxResp.AccessToken
}

func Run() {
	// return
	WXMgr.ReqAccessToken()
	for {
		select {
		case <-time.After(time.Second * 7200):
			WXMgr.ReqAccessToken()
		}
	}
}
