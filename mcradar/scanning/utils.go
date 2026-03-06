package scanning

import (
	"strconv"
	"strings"
)

const mask uint32 = 0xFF

func NumericIpToString(n uint32) string {
	segments := make([]string, 4)

	for i := 0; i < 4; i++ {
		segments[3-i] = strconv.FormatUint(uint64(mask&(n>>(i*8))), 10)
	}

	return strings.Join(segments, ".")
}

func IsReserved(addr string) bool {
	return strings.HasPrefix(addr, "0.") ||
		strings.HasPrefix(addr, "10.") ||
		strings.HasPrefix(addr, "127.") ||
		strings.HasPrefix(addr, "169.254.") ||
		strings.HasPrefix(addr, "192.0.0.") ||
		strings.HasPrefix(addr, "192.0.2.") ||
		strings.HasPrefix(addr, "192.88.99.") ||
		strings.HasPrefix(addr, "192.168.") ||
		strings.HasPrefix(addr, "198.51.100.") ||
		strings.HasPrefix(addr, "203.0.113.") ||
		strings.HasPrefix(addr, "233.252.0.")
}
