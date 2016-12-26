package auth

import (
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"time"
)

type AuthObject struct {
	Uid     int64
	Phone   string
	Openid  string
	Expires int64
	Atime   int64
}

type AuthService struct {
	app.Service
	Init   *app.InitTask
	Auth   *AuthTask
	Set    *AuthSetTask
	Remove *AuthRemoveTask

	Expires int64

	dispatch *kk.Dispatch
	objects  map[string]*AuthObject
}

func (S *AuthService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *AuthService) HandleInitTask(a *AuthApp, task *app.InitTask) error {

	S.objects = map[string]*AuthObject{}
	S.dispatch = kk.NewDispatch()

	var fn func() = nil

	fn = func() {

		var keys []string = []string{}
		var now = time.Now().Unix()

		for key, value := range S.objects {

			if value.Atime+value.Expires < now {
				keys = append(keys, key)
			}

		}

		for _, key := range keys {
			delete(S.objects, key)
		}

		S.dispatch.AsyncDelay(fn, time.Second*6)
	}

	S.dispatch.AsyncDelay(fn, time.Second*6)

	return nil
}

func (S *AuthService) HandleAuthTask(a *AuthApp, task *AuthTask) error {

	if task.Code == "" {
		task.Result.Errno = ERROR_AUTH_NOT_FOUND_CODE
		task.Result.Errmsg = "Not Found code"
		return nil
	}

	S.dispatch.Sync(func() {

		v, ok := S.objects[task.Code]

		if ok {
			v.Atime = time.Now().Unix()
			task.Result.Uid = v.Uid
			task.Result.Phone = v.Phone
			task.Result.Openid = v.Openid
		} else {
			task.Result.Errno = ERROR_AUTH_NOPERMISSION
			task.Result.Errmsg = "No Premission"
		}

	})

	return nil
}

func (S *AuthService) HandleAuthSetTask(a *AuthApp, task *AuthSetTask) error {

	if task.Code == "" {
		task.Result.Errno = ERROR_AUTH_NOT_FOUND_CODE
		task.Result.Errmsg = "Not Found code"
		return nil
	}

	S.dispatch.Sync(func() {

		v, ok := S.objects[task.Code]

		if ok {
			if task.Uid != 0 {
				v.Uid = task.Uid
			}
			if task.Phone != "" {
				v.Phone = task.Phone
			}
			if task.Openid != "" {
				v.Openid = task.Openid
			}
			v.Atime = time.Now().Unix()
		} else {
			v = &AuthObject{task.Uid, task.Phone, task.Openid, int64(time.Second) * S.Expires, time.Now().Unix()}
			S.objects[task.Code] = v
		}

		task.Result.Uid = v.Uid
		task.Result.Phone = v.Phone
		task.Result.Openid = v.Openid

	})

	return nil
}

func (S *AuthService) HandleAuthRemoveTask(a *AuthApp, task *AuthRemoveTask) error {

	if task.Code == "" {
		task.Result.Errno = ERROR_AUTH_NOT_FOUND_CODE
		task.Result.Errmsg = "Not Found code"
		return nil
	}

	S.dispatch.Sync(func() {
		delete(S.objects, task.Code)
	})

	return nil
}
