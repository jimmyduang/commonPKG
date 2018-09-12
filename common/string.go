package common

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
* 取指定长度的字符串
* @param string str 需要截取的字符串
* @param int start 开始位置(索引，从0开始)
* @param int length 截取的长度
* @return string 返回截取后的字符串
 */
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

/**
* 指定位置插入字符串
* @param string str 原始字符串
* @param string inser 需要传入的字符串
* @param int start 插入的位置
* @return  string 返回插入后的字符串
 */
func InsertStr(str, inser string, start int) string {
	res := str
	strLen := len(str)
	if start > 0 {
		if start == 0 {
			res = inser + res
		} else if strLen > start {
			pre := Substr(str, 0, start)
			suf := Substr(str, start, strLen-start)
			res = pre + inser + suf
		} else {
			res = res + inser
		}
	}

	return res
}

/**
* interface类型转成字符串
* @param interface{} 转换的值
* @return string 返回的字符串
 */
func InterfaceToString(inter interface{}) string {
	res := ""
	switch inter := inter.(type) {
	case bool:
		res = fmt.Sprintf("%t", inter)
	case int:
		res = fmt.Sprintf("%d", inter)
	case int64:
		res = fmt.Sprintf("%d", inter)
	case float64:
		res = strconv.FormatFloat(inter, 'f', -1, 64)
	case byte:
		res = fmt.Sprintf("%b", inter)
	case string:
		res = fmt.Sprintf("%s", inter)
	case *bool:
		res = fmt.Sprintf("%p", inter)
	case *int:
		res = fmt.Sprintf("%p", inter)
	case *int64:
		res = fmt.Sprintf("%p", inter)
	case *float64:
		res = fmt.Sprintf("%p", inter)
	case *string:
		res = fmt.Sprintf("%p", inter)
	}
	return res
}

/**
* byte转字符串
* @param byte 需要转换的byte数组
* @return string 返回的字符串
 */
func ByteToString(b []byte) string {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return string(b[0:i])
		}
	}
	return string(b)
}

func BankStr(str string) string {
	str_len := len(str)
	new_str := string(str[(str_len - 4):str_len])
	return "***************" + new_str
}

/*
 * 返回手机/邮箱等
 * 这种格式 135****1234
 */
func PrivacyInfoBasic(str string) string {

	new_str := ""
	str_len := len(str)

	salt := "****"
	salt_len := len(salt)

	if str_len > 2 {

		var str_c float64
		str_c = float64(str_len) / float64(salt_len)

		f_str_int_s := math.Ceil(str_c) //向上取整
		str_int_s := int(f_str_int_s)

		//f_str_int_x := math.Floor(str_c) //向下取整
		//str_int_x := int(f_str_int_x)

		str_prefix := string(str[0:str_int_s])
		str_suffix := string(str[(str_len - 1):str_len])
		if (salt_len + str_int_s) < str_len {
			str_suffix = string(str[(salt_len + str_int_s):str_len])
		}

		new_str = str_prefix + salt + str_suffix
	} else {

		new_str = salt
	}
	return new_str
}

/*
 * 隐私信息保护
 * 手机/邮箱/QQ
 */
func PrivacyInfo(str string) string {
	new_str := ""

	str_len := len(str)

	if str_len > 0 {

		is_email := strings.Contains(str, "@")
		if is_email {

			new_str_arr := strings.Split(str, "@")
			if len(new_str_arr) > 1 {

				new_str_prefix := PrivacyInfoBasic(new_str_arr[0])
				new_str_suffix := new_str_arr[1]

				new_str = new_str_prefix + "@" + new_str_suffix
			}

		} else {

			new_str = PrivacyInfoBasic(str)
		}
	}

	return new_str
}

/*
去除字符串右边的符号
@str 字符串
@tt 想删除的符号
*/
func Rtirm(str, tt string) string {
	new_str := strings.TrimRight(str, tt)
	return new_str
}
