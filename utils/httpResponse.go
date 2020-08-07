package utils

import logger "github.com/tal-tech/loggerX"

// Response the unified json structure
type response struct {
	Code    int         `json:"code"`
	Stat    int         `json:"stat"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func Success(v interface{}) interface{} {
	ret := response{Stat: 1, Code: 0, Message: "ok", Data: v}
	return ret
}

func Error(err error) interface{} {
	e := logger.NewError(err)
	ret := response{Stat: 0, Code: e.Code, Message: e.Message, Data: e.Info}
	return ret
}
