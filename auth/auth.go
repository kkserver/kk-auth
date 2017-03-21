package auth

import (
	"database/sql"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

type Auth struct {
	Id int64 `json:"id"`

	Code string `json:"code"`

	Uid      int64  `json:"uid,string,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Openid   string `json:"openid,omitempty"`
	DeviceId string `json:"deviceid,omitempty"`
	Expires  int64  `json:"expires,omitempty"`
	Atime    int64  `json:"atime"`

	Ctime int64 `json:"ctime"`
}

type IAuthApp interface {
	app.IApp
	GetDB() (*sql.DB, error)
	GetPrefix() string
	GetAuthTable() *kk.DBTable
	GetCacheExpires() int64
}

type AuthApp struct {
	app.App

	DB *app.DBConfig

	Remote    *remote.Service
	Auth      *AuthService
	AuthTable kk.DBTable
}

func (C *AuthApp) GetDB() (*sql.DB, error) {
	return C.DB.Get(C)
}

func (C *AuthApp) GetPrefix() string {
	return C.DB.Prefix
}

func (C *AuthApp) GetAuthTable() *kk.DBTable {
	return &C.AuthTable
}
