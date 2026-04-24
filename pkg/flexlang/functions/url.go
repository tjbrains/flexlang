// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"net/url"
	"strconv"
)

type URL struct {
	Host   string            `json:"host" expr:"host"`
	Query  string            `json:"query" expr:"query"`
	Port   int               `json:"port" expr:"port"`
	Path   string            `json:"path" expr:"path"`
	Hash   string            `json:"hash" expr:"hash"`
	Scheme string            `json:"scheme" expr:"scheme"`
	Opaque string            `json:"opaque" expr:"opaque"`
	User   map[string]string `json:"user" expr:"user"`
}

func NewURL(urlString string) URL {
	u, err := url.Parse(urlString)
	if err != nil {
		return URL{}
	}
	var user = map[string]string{}
	if u.User != nil {
		password, _ := u.User.Password()
		user = map[string]string{
			"username": u.User.Username(),
			"password": password,
		}
	}

	port, _ := strconv.Atoi(u.Port())

	return URL{
		Host:   u.Host,
		Query:  u.RawQuery,
		Port:   port,
		Path:   u.Path,
		Hash:   u.Fragment,
		Scheme: u.Scheme,
		Opaque: u.Opaque,
		User:   user,
	}
}
