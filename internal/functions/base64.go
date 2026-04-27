// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"encoding/base64"
)

type Base64Functions struct{}

func (this Base64Functions) Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (this Base64Functions) Decode(s string) string {
	if len(s) == 0 {
		return ""
	}
	result, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(result)
}
