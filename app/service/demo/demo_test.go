package demo

import (
	"context"
	"testing"
)

func Test_DoFun(t *testing.T) {
	ret, err := DoFun(context.Background(), "")
	if err != nil {
		t.Errorf("demoservice.DoFun test failed,err:%v", err)
	} else {
		t.Logf("demoservice.DoFun return %v", ret)
	}
}
