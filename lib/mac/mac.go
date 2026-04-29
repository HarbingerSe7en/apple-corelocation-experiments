package mac

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)

func Decode(i int64) string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", b[2], b[3], b[4], b[5], b[6], b[7])
}

func Encode(mac string) (int64, error) {
	// Remove colons from the MAC address
	macHex := strings.ReplaceAll(mac, ":", "")

	// Ensure the MAC address is valid
	if len(macHex) != 12 {
		return 0, fmt.Errorf("invalid MAC address length")
	}

	// Convert hexadecimal string to int64
	b, err := hex.DecodeString(macHex)
	if err != nil {
		return 0, fmt.Errorf("invalid hex")
	}
	// Pad the start with 0s
	if len(b) != 8 {
		for i := 0; i <= 8-len(b); i++ {
			b = append([]byte{0}, b...)
		}
	}
	return int64(binary.BigEndian.Uint64(b)), nil
}

func BytesToInt64(mac []byte) int64 {
	var i int64
	for _, b := range mac {
		i = i<<8 + int64(b)
	}
	return i
}
