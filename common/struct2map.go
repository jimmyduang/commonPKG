package common

import (
	"reflect"
	"strings"
)

//结构体转为map
func Struct2Map(obj interface{}, notcol string) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if !strings.Contains(notcol, t.Field(i).Name) {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}

	}
	return data
}
