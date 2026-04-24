// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
)

type CryptoHMACHash struct {
	hash hash.Hash

	Update func(data string) bool `expr:"update"`
	Sum    func() string          `expr:"sum"`
}

func NewCryptoHMAC(algorithm string, key string) (CryptoHMACHash, error) {
	var h hash.Hash
	switch algorithm {
	case "md5":
		h = hmac.New(md5.New, []byte(key))
	case "sha1":
		h = hmac.New(sha1.New, []byte(key))
	case "sha256":
		h = hmac.New(sha256.New, []byte(key))
	case "sha512":
		h = hmac.New(sha512.New, []byte(key))
	default:
		return CryptoHMACHash{}, errors.New("invalid hash algorithm")
	}

	return CryptoHMACHash{
		hash: h,
		Update: func(data string) bool {
			h.Reset()
			h.Write([]byte(data))
			return true
		},
		Sum: func() string {
			var data = h.Sum(nil)
			return hex.EncodeToString(data)
		},
	}, nil
}
