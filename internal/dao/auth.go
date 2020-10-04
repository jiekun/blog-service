// @Author: 2014BDuck
// @Date: 2020/10/4

package dao

import "github.com/blog-service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
