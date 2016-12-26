package auth

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type AuthRemoveResult struct {
	app.Result
}

type AuthRemoveTask struct {
	app.Task
	Code   string `json:"code"`
	Result AuthRemoveResult
}

func (T *AuthRemoveTask) GetResult() interface{} {
	return &T.Result
}
