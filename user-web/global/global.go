package global

import (
	"github.com/ShiyuCheng2018/mxshop-api/user-web/config"
	ut "github.com/go-playground/universal-translator"
)

var (
	Trans        ut.Translator
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)
