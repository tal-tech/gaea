package demoservice

import (
	"context"
	"math/rand"

	logger "github.com/tal-tech/loggerX"
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
