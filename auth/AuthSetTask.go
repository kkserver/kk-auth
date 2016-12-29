package auth

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type AuthSetResult struct {
	app.Result
	Uid      int64  `json:"uid,string,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Openid   string `json:"openid,omitempty"`
	DeviceId string `json:"deviceid,omitempty"`
}

type AuthSetTask struct {
	app.Task
	Code     string `json:"code"`
	Uid      int64  `json:"uid,string,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Openid   string `json:"openid,omitempty"`
	DeviceId string `json:"deviceid,omitempty"`
	Expires  int64  `json:"expires,omitempty"`
	Result   AuthSetResult
}

func (T *AuthSetTask) GetResult() interface{} {
	return &T.Result
}
