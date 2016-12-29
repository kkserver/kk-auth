package auth

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type AuthClearResult struct {
	app.Result
}

type AuthClearTask struct {
	app.Task
	Uid      int64  `json:"uid,string,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Openid   string `json:"openid,omitempty"`
	DeviceId string `json:"deviceid,omitempty"`
	Result   AuthRemoveResult
}

func (T *AuthClearTask) GetResult() interface{} {
	return &T.Result
}
