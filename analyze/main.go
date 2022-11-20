package analyze

import (
	"encoding/json"
	"fmt"

	"github.com/google/gopacket"
)

type Layer struct {
	Type string
	Data gopacket.Layer
}

type Packet struct {
	Metadata *gopacket.PacketMetadata
	Data     []byte
	Layers   []Layer
}

// Basic packet analyzer
func Print(packet gopacket.Packet) error {
	var layers []Layer

	for _, layer := range packet.Layers() {
		layers = append(layers, Layer{
			Type: layer.LayerType().String(),
			Data: layer,
		})
	}

	p := &Packet{
		Metadata: packet.Metadata(),
		Data:     packet.Data(),
		Layers:   layers,
	}

	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(json))

	return nil
}
