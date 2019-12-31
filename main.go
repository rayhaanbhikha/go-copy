package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/rayhaanbhikha/s3clip/s3"
	"github.com/urfave/cli/v2"
)

func main() {
	//TODO: validate access_id and secret_key exist.
	app := cli.NewApp()
	app.Name = "s3clip"
	app.Usage = "copy and paste from an s3 bucket"

	app.Commands = []*cli.Command{
		{
			Name:    "copy",
			Aliases: []string{"c"},
			Usage:   "copy to s3",
			Action:  copyCommand,
		},
		{
			Name:    "paste",
			Aliases: []string{"p"},
			Usage:   "paste from s3",
			Action: func(c *cli.Context) error {
				err := s3.Download()
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func copyCommand(c *cli.Context) error {
	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		return errors.New("No input provided")
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	return s3.Upload(string(data))
}
