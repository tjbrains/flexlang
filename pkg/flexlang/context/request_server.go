// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context

type RequestServerInfo struct {
	// 网站ID
	//
	// @expr id
	Id int64 `expr:"id" json:"id"`
}
