// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package utils

import (
	"os"
	"path/filepath"
)

func RootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "./"
	}

	var parent = dir
	var count = 32
	for {
		var modFile = parent + "/" + "go.mod"
		_, err = os.Stat(modFile)
		if err == nil {
			return parent
		}

		parent = filepath.Dir(parent)
		count--
		if count == 0 {
			return "./"
		}
	}
}
