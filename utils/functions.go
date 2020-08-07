package utils

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"time"

	"fmt"
	"math"

	"context"
	"github.com/spf13/cast"
	logger "github.com/tal-tech/loggerX"
	"github.com/tal-tech/xtools/jsutil"
)

func RemoveDuplicateAndEmpty(arr []string) []string {
	m := make(map[string]bool, 0)
	result := make([]string, 0)
	for _, v := range arr {
		if strings.Trim(v, " ") != "" {
			if _, ok := m[v]; !ok {
				m[v] = true
				result = append(result, v)
			}
		}
	}
	return result
}

func Convert(oldType interface{}, newType interface{}) (err error) {
	logger.D("[Convert]", "params:%v,%v, types:%v, %v", oldType, newType, reflect.TypeOf(oldType), reflect.TypeOf(newType))
	str, err := jsutil.Json.Marshal(oldType)
	if err != nil {
		return err
	}
	err = jsutil.Json.Unmarshal(str, &newType)
	if err != nil {
		return err
	}

	logger.D("[Convert]", "result:%v", newType)
	return nil
}

func ConvertWithMarshaledInput(oldStr string, newType interface{}) (back interface{}, err error) {
	err = jsutil.Json.Unmarshal([]byte(oldStr), &newType)
	if err != nil {
		return nil, err
	}
	return newType, nil
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
func MapUnion_Interface(dst *map[string]interface{}, src map[string]interface{}) {
	for index, value := range src {
		if _, ok := (*dst)[index]; !ok {
			(*dst)[index] = value
		}
	}
}

func MapUnion_Map(dst *map[string]map[string]interface{}, src map[string]map[string]interface{}) {
	for index, value := range src {
		if _, ok := (*dst)[index]; !ok {
			(*dst)[index] = value
		}
	}
}

func MapUnion_Int64Slice(dst *map[string][]int64, src map[string][]int64) {
	for index, value := range src {
		if _, ok := (*dst)[index]; !ok {
			(*dst)[index] = value
		}
	}
}
func MapUnion_StringSlice(dst *map[string][]string, src map[string][]string) {
	for index, value := range src {
		if _, ok := (*dst)[index]; !ok {
			(*dst)[index] = value
		}
	}
}

//todo: check entry existences
func SliceUnion_Int64(dst *[]int64, src []int64) {
	for _, value := range src {
		*dst = append(*dst, value)
	}
}

func Int64SliceToInterfaceSlice(params []int64) []interface{} {
	logger.D("[Int64SliceToInterfaceSlice]", "params:%v", params)

	result := make([]interface{}, 0)
	for _, element := range params {
		result = append(result, element)
	}

	logger.D("[Int64Slicetointerfaceslice]", "result:%v", result)
	return result
}
func StringSliceToInterfaceSlice(params []string) []interface{} {
	logger.D("[StringSliceToInterfaceSlice]", "params:%v", params)

	result := make([]interface{}, 0)
	for _, element := range params {
		result = append(result, element)
	}

	logger.D("[Stringslicetointerfaceslice]", "result:%v", result)
	return result
}

func TimeCostDuration(tag string, start time.Time) {
	du := time.Now().Sub(start)
	logger.D(tag, "Cost %d", du.Nanoseconds())
}

func JsonMustMarshal(v interface{}) string {
	result, err := jsutil.Json.Marshal(v)
	if err != nil {
		panic(logger.NewError("JsonMustMarshal failed, error:" + err.Error()))
	}

	return string(result)
}

/***
remove comment and whitespace
***/
func JsonMinify(str string) string {
	in_string := false
	in_single_comment := false
	in_multi_comment := false
	string_opener := string('x')

	var retStr string = ""

	size := len(str)
	for i := 0; i < size; i++ {

		c := fmt.Sprintf("%s", []byte{str[i]})
		next := math.Min(cast.ToFloat64(i+2), cast.ToFloat64(size))
		cc := fmt.Sprintf("%s", []byte(str[i:cast.ToInt64(next)]))

		if in_string {
			if c == string_opener {
				in_string = false
				retStr += c
			} else if c == "\\" {
				retStr += cc
				i++
			} else {
				retStr += c
			}
		} else if in_single_comment {
			if c == "\r" || c == "\n" {
				in_single_comment = false
			}
		} else if in_multi_comment {
			if cc == "*/" {
				in_multi_comment = false
				i++
			}
		} else {
			if cc == "/*" {
				in_multi_comment = true
				i++
			} else if cc == "//" {
				in_single_comment = true
				i++
			} else if c[0] == '"' || c[0] == '\'' {
				in_string = true
				string_opener = c
				retStr += c
			} else if c != " " && c != "\t" && c != "\n" && c != "\r" {
				retStr += c
			}
		}
	}
	return retStr
}

func MapIntersect(s1, s2 []string) []string {
	length1 := len(s1)
	length2 := len(s2)
	intersect := make([]string, 0)
	s1map := make(map[string]string, length1)
	for i := 0; i < length1; i++ {
		s1map[s1[i]] = s1[i]
	}

	for j := 0; j < length2; j++ {
		if _, ok := s1map[s2[j]]; ok {
			intersect = append(intersect, s1map[s2[j]])
		}
	}
	return intersect
}

func TransferToContext(c *gin.Context) context.Context {
	ctx := context.Background()
	for k, v := range c.Keys {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}
