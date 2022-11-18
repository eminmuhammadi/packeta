package analyze

import (
	"fmt"

	"github.com/google/gopacket"
)

// Basic packet analyzer
func Print(packet gopacket.Packet) error {
	fmt.Println(packet)

	return nil
}
