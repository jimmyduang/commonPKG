package common

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

/**
* 整形转byte数组
 */
func IntToByte(data int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((data>>shift)&mask))
	}
	return result
}

/**
* byte数组转int64
 */
func ByteToInt(data []byte) int64 {
	b_buf := bytes.NewBuffer(data)
	var res int32
	binary.Read(b_buf, binary.BigEndian, &res)
	return int64(res)
}

func ByteToUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

//字符串转int64
func Str2Int64(str string) (int64, error) {
	i64, err := strconv.ParseInt(str, 10, 64)
	return i64, err
}

//字符串转int
func Str2Int(str string) (int, error) {
	i64, err := Str2Int64(str)
	return int(i64), err
}
