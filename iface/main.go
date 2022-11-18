package iface

import (
	"errors"
	"fmt"

	"github.com/google/gopacket/pcap"
)

// Returns a list of all interfaces on the system.
func GetIfaces() ([]pcap.Interface, error) {
	return pcap.FindAllDevs()
}

// Returns the interface with the given name.
func SelectInterfaceByName(name string) (pcap.Interface, error) {
	ifaces, err := GetIfaces()
	if err != nil {
		return pcap.Interface{}, err
	}

	for _, iface := range ifaces {
		if iface.Name == name {
			return iface, nil
		}
	}

	return pcap.Interface{}, errors.New("interface not found")
}

// Returns the interface with the given description.
func Print() error {
	ifaces, err := GetIfaces()
	if err != nil {
		return err
	}

	table, err := GenerateTable(ifaces)
	if err != nil {
		return err
	}

	fmt.Println(table)

	return nil
}
