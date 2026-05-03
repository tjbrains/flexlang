// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"net/netip"

	"github.com/tjbrains/flexlang/internal/utils"
)

type NetIPFunctions struct{}

func (this NetIPFunctions) IsValid(ip string) bool {
	_, err := netip.ParseAddr(ip)
	return err == nil
}

func (this NetIPFunctions) IsIPv4(ip string) bool {
	addr, err := netip.ParseAddr(ip)
	return err == nil && addr.Is4()
}

func (this NetIPFunctions) IsIPv6(ip string) bool {
	addr, err := netip.ParseAddr(ip)
	return err == nil && addr.Is6()
}

func (this NetIPFunctions) IsBetween(ip string, startIP string, endIP string) bool {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false
	}
	startAddr, err := netip.ParseAddr(startIP)
	if err != nil {
		return false
	}
	endAddr, err := netip.ParseAddr(endIP)
	if err != nil {
		return false
	}
	return addr.Compare(startAddr) >= 0 && addr.Compare(endAddr) <= 0
}

func (this NetIPFunctions) IsInCIDR(ip string, cidr string) bool {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false
	}
	cidrAddr, err := netip.ParsePrefix(cidr)
	if err != nil {
		return false
	}
	return cidrAddr.Contains(addr)
}

func (this NetIPFunctions) IsInRanges(ip string, ranges []any) bool {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false
	}
	for _, r := range ranges {
		var r2 = r.([]any)
		if len(r2) != 2 {
			continue
		}
		var start = utils.String(r2[0])
		var end = utils.String(r2[1])
		startAddr, err := netip.ParseAddr(start)
		if err != nil {
			return false
		}
		endAddr, err := netip.ParseAddr(end)
		if err != nil {
			return false
		}
		if addr.Compare(startAddr) >= 0 && addr.Compare(endAddr) <= 0 {
			return true
		}
	}
	return false
}
