package cmd

import (
	analyze "github.com/eminmuhammadi/packeta/analyze"
	capture "github.com/eminmuhammadi/packeta/capture"
	cli "github.com/urfave/cli/v2"
)

func Capture() *cli.Command {
	return &cli.Command{
		Name:  "capture",
		Usage: "Captures network packets",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "interface",
				Aliases:  []string{"i"},
				Usage:    "Input file",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "filter",
				Aliases:  []string{"f"},
				Usage:    "Filter expression",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			ifaceName := ctx.String("interface")
			filter := ctx.String("filter")

			// Get the handler
			handler, err := capture.Handler(ifaceName)
			if err != nil {
				return err
			}

			defer handler.Close()

			// Filter the packets
			if filter != "" {
				handler.Filter(filter)
			}

			// Start capturing
			handler.Start(analyze.Print)

			return err
		},
	}
}
