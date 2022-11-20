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

func PrettyPrint(packet gopacket.Packet) error {
	return Analyzer(packet, true)
}

func Print(packet gopacket.Packet) error {
	return Analyzer(packet, false)
}

// Basic packet analyzer
func Analyzer(packet gopacket.Packet, pretty bool) error {
	var layers []Layer
	var jsonByte []byte
	var err error

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

	if pretty {
		jsonByte, err = json.MarshalIndent(p, "", "  ")
	} else {
		jsonByte, err = json.Marshal(p)
	}

	if err != nil {
		return err
	}

	fmt.Println(string(jsonByte))

	return nil
}
