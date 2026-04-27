// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import (
	"errors"
	"strings"

	"github.com/expr-lang/expr/file"
)

func FixError(code string, err error) (string, bool) {
	if err == nil {
		return code, false
	}

	fileErr, ok := errors.AsType[*file.Error](err)
	if !ok {
		return code, false
	}

	if fileErr.Message == "unexpected token EOF" {
		var loc = fileErr.Location
		var r = []rune(code)
		if loc.From >= 0 && loc.To <= len(r) {
			var snippet = string(r[loc.From:loc.To])
			if strings.TrimSpace(snippet) == "}" {
				code += "\n else { false }"
			}
		}
	}

	return code, true
}

func TrimError(err error) error {
	if err == nil {
		return nil
	}

	fileErr, ok := errors.AsType[*file.Error](err)
	if !ok {
		return err
	}

	if len(fileErr.Message) > 300 {
		fileErr.Message = fileErr.Message[:100] + " [...] " + fileErr.Message[len(fileErr.Message)-100:]
	}

	return fileErr
}
