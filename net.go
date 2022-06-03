package utils

import "net"

/**
* 아이피 주소 검증
**/
func IsValidIP(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	}
	return true
}
