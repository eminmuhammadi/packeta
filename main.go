package main

import (
	"fmt"
	"log"
	"os"

	cmd "github.com/eminmuhammadi/packeta/cmd"
	cli "github.com/urfave/cli/v2"
)

const (
	VERSION    = "0.0.0"
	BUILD_ID   = ""
	BUILD_TIME = ""
)

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
		Copyright: "github.com/eminmuhammadi/packeta",
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
