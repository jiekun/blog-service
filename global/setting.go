// @Author: 2014BDuck
// @Date: 2020/9/13

package global

import (
	"github.com/blog-service/pkg/logger"
	"github.com/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	Logger          *logger.Logger
)
