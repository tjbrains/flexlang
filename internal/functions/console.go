// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"encoding/json"
	"fmt"
)

type ConsoleFunctions struct {
}

func (this ConsoleFunctions) Log(printer func(s ...string), data ...any) bool {
	var l = len(data)
	if l == 0 {
		return true
	}
	var args = make([]string, 0, l)
	for _, d := range data {
		switch d2 := d.(type) {
		case string:
			args = append(args, d2)
		case bool:
			if d2 {
				args = append(args, "true")
			} else {
				args = append(args, "false")
			}
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			args = append(args, fmt.Sprintf("%d", d2))
		case float32, float64:
			args = append(args, fmt.Sprintf("%g", d2))
		case []byte:
			args = append(args, string(d2))
		case []rune:
			args = append(args, string(d2))
		default:
			tJSON, _ := json.Marshal(d2)
			args = append(args, string(tJSON))
		}
	}

	printer(args...)

	return true
}
