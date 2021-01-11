package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/gucchisk/bytestring"
	"github.com/gucchisk/byconv/flags"
)

func main() {
	var inputFormat, outputFormat string
	app := &cli.App{
		Name: "byconv",
		UsageText: "byconv [options]",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "input",
				Usage: "input",
				Aliases: []string{"i"},
				Destination: &inputFormat,
			},
			&cli.StringFlag{
				Name: "output",
				Usage: "output",
				Aliases: []string{"o"},
				Destination: &outputFormat,
			},
		},
		Action: func(c *cli.Context) error {
			if inputFormat != "" {
				fmt.Printf("input: %s\n", inputFormat)
			}

			filename := c.Args().Get(0)
			var reader io.Reader
			deferFunc := func() error {
				return nil
			}
			if filename == "" {
				reader = os.Stdin
			} else {
				file, err := os.Open(filename)
				if err != nil {
					return err
				}
				deferFunc = file.Close
				reader = file
			}
			defer deferFunc()
			
			byteArray, err := ioutil.ReadAll(reader)
			if err != nil {
				return err
			}

			input, err := flags.NewFormat(inputFormat)
			if err != nil {
				return err
			}
			bytes, err := bytestring.NewBytes(byteArray, bytestring.Type(input))
			if err != nil {
				return err
			}
			output, err := flags.NewFormat(outputFormat)
			if err != nil {
				return err
			}

			var result string
			switch output {
			case bytestring.Ascii:
				result = bytes.String()
			case bytestring.Hex:
				result = bytes.HexString()
			case bytestring.Base64:
				result = bytes.Base64()
			default:
				return fmt.Errorf("error")
			}
			fmt.Printf("%s", result)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
