package auth

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type AuthCreateResult struct {
	app.Result
	Code string `json:"code,omitempty"`
}

type AuthCreateTask struct {
	app.Task
	Uid      int64  `json:"uid,string,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Openid   string `json:"openid,omitempty"`
	DeviceId string `json:"deviceid,omitempty"`
	Expires  int64  `json:"expires,omitempty"`
	Result   AuthCreateResult
}

func (T *AuthCreateTask) GetResult() interface{} {
	return &T.Result
}
