package main

import (
	"encoding/json"
	"fmt"
	"gaea/app/router"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	logger "git.100tal.com/wangxiao_go_lib/xesLogger"
	"git.100tal.com/wangxiao_go_lib/xesServer/ginhttp"
	"git.100tal.com/wangxiao_go_lib/xesTools/confutil"
	"git.100tal.com/wangxiao_go_lib/xesTools/flagutil"
	"github.com/spf13/cast"
)

func TestServer(t *testing.T) {
	flagutil.SetConfig("conf/conf.ini")
	confutil.SetConfPathPrefix(os.Getenv("GOPATH") + "/src/gaea")
	logger.InitLogger(os.Getenv("GOPATH") + "/src/gaea/conf/log.xml")
	confutil.InitConfig()
	//signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	s := ginhttp.NewServer()
	router.RegisterRouter(s.GetGinEngine())
	s.AddServerBeforeFunc(s.InitConfig())
	defer logger.Close()

	go func() {
		err := s.Serve()
		if err != nil {
			t.Fatalf("TestServer start failed,err:%v", err)
		}
	}()
	time.Sleep(1 * time.Second)
	send(t)
	/*
		<-exit
		s.Stop()
	*/
}

func send(t *testing.T) {
	t.Log("sending")
	httpclient := http.DefaultClient
	var resp *http.Response
	var req *http.Request
	port := confutil.GetConf("Server", "addr")
	req, err := http.NewRequest("GET", "http://127.0.0.1"+port+"/demo/test", nil)
	if err != nil {
		t.Fatalf("TestServer new request failed,err:%v", err)
	}
	resp, err = httpclient.Do(req)
	if err != nil {
		t.Fatalf("TestServer do request failed,err:%v", err)
	}

	if resp.StatusCode == 200 {
		ret, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("TestServer parse response failed,err:%v", err)
		}
		result := make(map[string]interface{}, 0)
		fmt.Println(string(ret))
		if err = json.Unmarshal(ret, &result); err != nil {
			t.Fatalf("TestServer parse response failed,err:%v", err)
		} else {
			code := cast.ToInt(result["code"])
			if code != 0 {
				t.Fatalf("TestServer response failed,code:%v", code)
			}
		}
		resp.Body.Close()
	} else {
		t.Fatalf("TestServer response failed,http code type err, code:%v", resp.StatusCode)
	}

	pid := os.Getpid()
	process, err := os.FindProcess(pid)
	if err != nil {
		t.Fatalf("TestServer failed,err:%v", err)
	}
	/*once := &sync.Once{}
	once.Do(func() {
		t.Log("restart")
		process.Signal(syscall.SIGUSR2)
		time.Sleep(10 * time.Second)
	})
	*/
	t.Log("close")
	process.Signal(syscall.SIGTERM)

	time.Sleep(1 * time.Second)
}
