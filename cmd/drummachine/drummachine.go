package main

import (
	"context"
	"github.com/IyadAssaf/go-drummachine/pkg/drummachine"
	"github.com/urfave/cli"
	"log"
	"os"
)


func main() {
	app := &cli.App{
		Name: "drummachine",
		Usage: "play those drums",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "kit",
				Value: "808",
				Usage: "808,909",
			},
		},
		Action: func(c *cli.Context) error {

			s, err := drummachine.NewDrumMachine(c.String("kit"))
			if err != nil {
				return err
			}
			defer s.Close()
			//TODO pass cli context
			return s.SetupInstrument(context.Background())
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}