// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package utils

import (
	"fmt"
	"strconv"
)

// String 将值转换成字符串
func String(value any) string {
	if value == nil {
		return ""
	}
	switch x := value.(type) {
	case int:
		return strconv.Itoa(x)
	case int64:
		return strconv.FormatInt(x, 10)
	case uint64:
		return strconv.FormatUint(x, 10)
	case int8, int16, int32, uint, uint8, uint16, uint32:
		return fmt.Sprintf("%d", x)
	case float32, float64:
		return fmt.Sprintf("%g", x)
	case []byte:
		return string(x)
	case []rune:
		return string(x)
	case string:
		return x
	}
	return fmt.Sprintf("%#v", value)
}
