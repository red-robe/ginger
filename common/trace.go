package common

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func GenTraceId(ip string) (traceId string) {
	now := time.Now()
	timestamp := uint32(now.Unix())
	timeNano := now.UnixNano()
	b := bytes.Buffer{}

	b.WriteString(fmt.Sprintf("%08x", timestamp&0xffffffff))
	b.WriteString(fmt.Sprintf("%04x", timeNano&0xffff))
	b.WriteString(fmt.Sprintf("%06x", rand.Int31n(1<<24)))

	netIP := net.ParseIP(ip)
	if netIP.To4() == nil {
		b.WriteString("00000000")
	} else {
		b.WriteString(hex.EncodeToString(netIP.To4()))
	}
	return b.String()
}
