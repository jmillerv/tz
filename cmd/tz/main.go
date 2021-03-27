package main

import (
	"fmt"
	tz "github.com/jmillerv/timezone/pkg/tz_converter"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main(){
	app := &cli.App{
		Name: "tz",
		Usage: "command line tool for retrieving and converting times from different timezones",
		Commands: []*cli.Command{
			{
				Name: "hello",
				Description: "say hello to tz",
				Action: func(c *cli.Context) error {
					fmt.Println("Hi!")
					return nil
				},
			},
			{
				Name: "CurrentTime",
				Usage: "ct <America/Boise>",
				Description: "returns current time based on IANA timezone name",
				Aliases: []string{"ct"},
				Action: func(c *cli.Context) error {
					location := c.Args().Get(0)
					if location == "" {
						location = "Local"
					}
					time, err := tz.TimeIn(time.Now(), location)
					if err != nil {
						return err
					}
					log.Printf("current time in %s: %s\n", location, time.Format("Mon Jan _2 15:04:05 2006"))
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}