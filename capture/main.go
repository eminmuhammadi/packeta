package capture

import (
	iface "github.com/eminmuhammadi/packeta/iface"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Sniffer is a function that analyzes the given packet.
type Sniffer func(packet gopacket.Packet) error

// Handle is a wrapper for pcap.Handle
type Handle pcap.Handle

// Returns the handler for the given interface
func Handler(name string) (*Handle, error) {
	// Get the interface
	iface, err := iface.SelectInterfaceByName(name)
	if err != nil {
		return nil, err
	}

	// Open the interface
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return nil, err
	}

	return (*Handle)(handle), nil
}

// Filter the packets
func (handler *Handle) Filter(expression string) error {
	return (*pcap.Handle)(handler).SetBPFFilter(expression)
}

// Closes the handler
func (handler *Handle) Close() {
	(*pcap.Handle)(handler).Close()
}

// Analyzes the given packet using the given analyzer.
func (handler *Handle) Start(sniffer Sniffer) {
	// Create a packet source
	packetSource := gopacket.NewPacketSource((*pcap.Handle)(handler), (*pcap.Handle)(handler).LinkType())

	// Loop through the packets
	for packet := range packetSource.Packets() {
		sniffer(packet)
	}
}
