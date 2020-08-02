package utils

import (
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	logger "git.100tal.com/wangxiao_go_lib/xesLogger"

	"git.100tal.com/wangxiao_go_lib/xesTools/confutil"
)

/*
先按url精确匹配,如果url精确匹配不到,再按uri分级去匹配
*/
func GetMockData(path string) []byte {
	u, err := url.Parse(path)
	if err != nil {
		logger.E("Parse Url Failed,Path:", path)
		return []byte{}
	}
	path = u.RequestURI()
	ret := GetConfData(path)
	if len(ret) == 0 {
		uris := strings.Split(path, "?")
		ret = GetConfData(uris[0])
		if len(ret) == 0 {
			uris := strings.Split(uris[0], "/")
			l := len(uris)
			for i := l - 1; i > 0; i-- {
				ret = GetConfData(strings.Join(uris[0:i], "/"))
				if len(ret) > 0 {
					return ret
				}
			}
		}
	}
	return ret
}

func GetConfData(url string) []byte {
	ret := make([]byte, 0)
	if path := confutil.GetConf("Mock", url); path != "" {
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			logger.E("Open File Failed,File:", path)
			return ret
		}
		ret, err = ioutil.ReadAll(file)
		if err != nil {
			logger.E("Read File Failed,File:", path)
			return ret
		}
	}
	return ret
}
