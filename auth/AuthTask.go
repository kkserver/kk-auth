package auth

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type AuthResult struct {
	app.Result
	Uid      int64  `json:"uid,string,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Openid   string `json:"openid,omitempty"`
	DeviceId string `json:"deviceid,omitempty"`
}

type AuthTask struct {
	app.Task
	Code   string `json:"code"`
	Result AuthResult
}

func (T *AuthTask) GetResult() interface{} {
	return &T.Result
}
