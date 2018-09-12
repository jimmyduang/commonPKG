package common

import (
	"strings"

	"github.com/go-xorm/xorm"
)

func Webcodecheck(adminWebCode, targetWebcode string, Orm *xorm.Engine) bool {
	istrue := true
	if targetWebcode == "all" {
		weblist := []map[string]string{}
		Orm.Table("web_list").Cols("web_code").Find(&weblist)
		for index, _ := range weblist {
			if !strings.Contains(adminWebCode, weblist[index]["web_list"]) {
				istrue = false
			}
		}
	} else {
		webcodes := strings.Split(targetWebcode, ",")
		for index, _ := range webcodes {
			if !strings.Contains(adminWebCode, webcodes[index]) {
				istrue = false
			}
		}
	}

	return istrue
}
