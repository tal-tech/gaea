package demoservice

import (
	"context"
	"math/rand"

	logger "git.100tal.com/wangxiao_go_lib/xesLogger"
)

func DoFun(ctx context.Context, param string) (ret map[string]interface{}, err error) {
	if rand.Intn(100)%2 == 0 {
		return nil, logger.NewError("dofun fail")
	} else {
		ret = make(map[string]interface{}, 2)
		ret["ret1"] = "dofun ok"
		ret["ret1"] = "Welcome to use gaea!"

	}
	return
}
