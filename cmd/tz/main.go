package main

import (
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

func main(){
	app := &cli.App{
		Name: "tz",
		Usage: "command line tool for retrieving and converting times from different timezones",
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}