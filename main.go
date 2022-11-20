package main

import (
	"fmt"
	"log"
	"os"

	cmd "github.com/eminmuhammadi/packeta/cmd"
	cli "github.com/urfave/cli/v2"
)

var VERSION = "0.1.1-dev"
var BUILD_ID = "0"
var BUILD_TIME = "0"

// Commands list
var Commands = []*cli.Command{
	cmd.Iface(),
	cmd.Capture(),
}

// Main entry point
func main() {
	fmt.Println(`
                            /                    
     ___     __     ___    /-__   ___   _/_-   __
    /   )  /   )   /   '  /(     /___)  /    /   )
   /___/  (___(_  (___  _/  \   (___   (_   (___(_
  /                                          
 /
 `)
	app := &cli.App{
		Name:      "packeta",
		Usage:     "Network packet analyzer",
		Version:   VERSION,
		Copyright: "packeta  Copyright (C) 2022  Emin Muhammadi",
		ExtraInfo: func() map[string]string {
			return map[string]string{
				"LICENSE":    "The GNU General Public License",
				"VERSION":    VERSION,
				"BUILD":      BUILD_ID,
				"BUILD_TIME": BUILD_TIME,
			}
		},
		Commands: Commands,
		Suggest:  true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
