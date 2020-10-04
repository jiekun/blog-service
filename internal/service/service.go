// @Author: 2014BDuck
// @Date: 2020/9/26

package service

import (
	"context"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/dao"
	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
