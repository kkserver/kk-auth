package auth

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type AuthResult struct {
	app.Result
	Auth *Auth `json:"auth,omitempty"`
}

type AuthTask struct {
	app.Task
	Code   string `json:"code"`
	Result AuthResult
}

func (T *AuthTask) GetResult() interface{} {
	return &T.Result
}
