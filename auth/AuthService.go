package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"math/rand"
	"time"
)

type AuthService struct {
	app.Service

	Get    *AuthTask
	Set    *AuthSetTask
	Create *AuthCreateTask
	Remove *AuthRemoveTask
}

func (S *AuthService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *AuthService) HandleAuthTask(a IAuthApp, task *AuthTask) error {

	if task.Code == "" {
		task.Result.Errno = ERROR_AUTH_NOT_FOUND_CODE
		task.Result.Errmsg = "Not Found code"
		return nil
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_AUTH
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Auth{}

	rows, err := kk.DBQuery(db, a.GetAuthTable(), a.GetPrefix(), " WHERE code=?", task.Code)

	if err != nil {
		task.Result.Errno = ERROR_AUTH
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {

		scanner := kk.NewDBScaner(&v)

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_AUTH
			task.Result.Errmsg = err.Error()
			return nil
		}

		if v.Atime+v.Expires >= time.Now().Unix() {

			task.Result.Auth = &v

			_, _ = kk.DBUpdateWithKeys(db, a.GetAuthTable(), a.GetPrefix(), &v, map[string]bool{"atime": true})

		} else {
			task.Result.Errno = ERROR_AUTH_NOPERMISSION
			task.Result.Errmsg = "no permission"
			return nil
		}

	} else {
		task.Result.Errno = ERROR_AUTH_NOPERMISSION
		task.Result.Errmsg = "no permission"
		return nil
	}

	return nil

}

func (S *AuthService) HandleAuthSetTask(a IAuthApp, task *AuthSetTask) error {

	if task.Code == "" {
		task.Result.Errno = ERROR_AUTH_NOT_FOUND_CODE
		task.Result.Errmsg = "Not Found code"
		return nil
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_AUTH
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Auth{}

	rows, err := kk.DBQuery(db, a.GetAuthTable(), a.GetPrefix(), " WHERE code=?", task.Code)

	if err != nil {
		task.Result.Errno = ERROR_AUTH
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {

		scanner := kk.NewDBScaner(&v)

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_AUTH
			task.Result.Errmsg = err.Error()
			return nil
		}

		if task.Uid != 0 {
			v.Uid = task.Uid
		}
		if task.Phone != "" {
			v.Phone = task.Phone
		}
		if task.Openid != "" {
			v.Openid = task.Openid
		}
		if task.DeviceId != "" {
			v.DeviceId = task.DeviceId
		}

		if task.Expires != 0 {
			v.Expires = task.Expires
		}

		v.Atime = time.Now().Unix()

		_, err = kk.DBUpdate(db, a.GetAuthTable(), a.GetPrefix(), &v)

		if err != nil {
			task.Result.Errno = ERROR_AUTH
			task.Result.Errmsg = err.Error()
			return nil
		}

		task.Result.Auth = &v

	} else {
		task.Result.Errno = ERROR_AUTH_NOPERMISSION
		task.Result.Errmsg = "no permission"
		return nil
	}

	return nil
}

func NewCode() string {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("auth %d %d M(&YGHJKL:", time.Now().UnixNano(), rand.Intn(100000))))
	return hex.EncodeToString(m.Sum(nil))
}

func (S *AuthService) HandleAuthCreateTask(a IAuthApp, task *AuthCreateTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_AUTH
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Auth{}
	v.Code = NewCode()
	v.Uid = task.Uid
	v.Phone = task.Phone
	v.DeviceId = task.DeviceId
	v.Openid = task.Openid
	v.Expires = task.Expires
	v.Atime = time.Now().Unix()
	v.Ctime = v.Atime

	_, err = kk.DBInsert(db, a.GetAuthTable(), a.GetPrefix(), &v)

	if err != nil {
		task.Result.Errno = ERROR_AUTH
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Auth = &v

	return nil
}

func (S *AuthService) HandleAuthRemoveTask(a IAuthApp, task *AuthRemoveTask) error {

	if task.Code == "" {
		task.Result.Errno = ERROR_AUTH_NOT_FOUND_CODE
		task.Result.Errmsg = "Not Found code"
		return nil
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_AUTH
		task.Result.Errmsg = err.Error()
		return nil
	}

	_, err = kk.DBDelete(db, a.GetAuthTable(), a.GetPrefix(), " WHERE code=?", task.Code)

	if err != nil {
		task.Result.Errno = ERROR_AUTH
		task.Result.Errmsg = err.Error()
		return nil
	}

	return nil
}
