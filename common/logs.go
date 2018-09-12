package common

import (
	"fmt"
)

func LogsWithcontent(arg ...string) {
	content := ""
	for _, v := range arg {
		content += v
		content += "-----"
	}
	fmt.Println(content)
}
