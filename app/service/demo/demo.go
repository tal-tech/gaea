package demo

import (
	"context"
)

func DoFun(ctx context.Context, param string) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{}, 2)
	ret["ret1"] = "dofun ok"
	ret["ret1"] = "Welcome to use gaea!"
	return
}
