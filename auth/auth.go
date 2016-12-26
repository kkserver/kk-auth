package auth

import (
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

type AuthApp struct {
	app.App
	Remote *remote.Service
	Auth   *AuthService
}
