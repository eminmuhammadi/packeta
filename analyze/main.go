package analyze

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/gopacket"
)

type Layer struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
	Payload  string `json:"payload"`
}

type Metadata struct {
	Timestamp      time.Time     `json:"timestamp"`
	CaptureLength  int           `json:"capture_length"`
	Length         int           `json:"length"`
	InterfaceIndex int           `json:"interface_index"`
	AncillaryData  []interface{} `json:"ancillary_data"`
	Truncated      bool          `json:"truncated"`
	Endpoints      Endpoint      `json:"endpoints"`
}

type Endpoint struct {
	NetworkLayer   Flow `json:"network_layer"`
	LinkLayer      Flow `json:"link_layer"`
	TransportLayer Flow `json:"transport_layer"`
}

type Flow struct {
	Type        string `json:"type"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type Packet struct {
	Layers   []Layer   `json:"layers"`
	Data     string    `json:"data"`
	Metadata *Metadata `json:"metadata"`
}

// Basic packet analyzer
func Print(packet gopacket.Packet) error {
	data, err := DataDecode(packet)
	if err != nil {
		return err
	}

	json, err := JSON(data)
	if err != nil {
		return err
	}

	fmt.Println(json)

	return nil
}

// Decode packet data
func DataDecode(packet gopacket.Packet) (*Packet, error) {
	var layers []Layer

	for _, layer := range packet.Layers() {
		layers = append(layers, Layer{
			Name:     layer.LayerType().String(),
			Contents: encode(layer.LayerContents()),
			Payload:  encode(layer.LayerPayload()),
		})
	}

	data := &Packet{
		Layers: layers,
		Data:   encode(packet.Data()),
		Metadata: &Metadata{
			Timestamp:      packet.Metadata().Timestamp,
			CaptureLength:  packet.Metadata().CaptureLength,
			Length:         packet.Metadata().Length,
			InterfaceIndex: packet.Metadata().InterfaceIndex,
			AncillaryData:  packet.Metadata().AncillaryData,
			Truncated:      packet.Metadata().Truncated,
			Endpoints: Endpoint{
				NetworkLayer: Flow{
					Type:        packet.NetworkLayer().NetworkFlow().EndpointType().String(),
					Source:      packet.NetworkLayer().NetworkFlow().Src().String(),
					Destination: packet.NetworkLayer().NetworkFlow().Dst().String(),
				},
				LinkLayer: Flow{
					Type:        packet.LinkLayer().LinkFlow().EndpointType().String(),
					Source:      packet.LinkLayer().LinkFlow().Src().String(),
					Destination: packet.LinkLayer().LinkFlow().Dst().String(),
				},
				TransportLayer: Flow{
					Type:        packet.TransportLayer().TransportFlow().EndpointType().String(),
					Source:      packet.TransportLayer().TransportFlow().Src().String(),
					Destination: packet.TransportLayer().TransportFlow().Dst().String(),
				},
			},
		},
	}

	return data, nil
}

// Json pretty print
func JSON(data *Packet) (string, error) {
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// Encoder
func encode(b []byte) string {
	// byte to base64
	return b64.StdEncoding.EncodeToString(b)
}
