package cmd

import (
	iface "github.com/eminmuhammadi/packeta/iface"
	cli "github.com/urfave/cli/v2"
)

func Iface() *cli.Command {
	return &cli.Command{
		Name:    "interfaces",
		Aliases: []string{"ifaces"},
		Usage:   "List all interfaces on the system",
		Flags:   []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			// Return the list of all interfaces on the system.
			iface.Print()

			return nil
		},
	}
}
