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
