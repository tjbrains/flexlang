// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import "encoding/json"

type JSONFunctions struct{}

func (this JSONFunctions) Parse(text string) any {
	var v any
	_ = json.Unmarshal([]byte(text), &v)
	return v
}

func (this JSONFunctions) Stringify(value any) string {
	data, _ := json.Marshal(value)
	return string(data)
}
