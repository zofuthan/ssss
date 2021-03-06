package ssss

import (
	"time"
)

type Config struct {
	HttpAddr     string
	HttpPort     int
	UseFcgi      bool
	RecoverPanic bool
	RunMode      int8 //0=prod，1=dev
	TemplatePath string
	ReadTimeout  time.Duration // maximum duration before timing out read of the request, 默认:5*time.Second(5秒超时)
	WriteTimeout time.Duration // maximum duration before timing out write of the response, 默认:0(不超时)
}

const RUNMODE_PROD int8 = 0
const RUNMODE_DEV int8 = 1
