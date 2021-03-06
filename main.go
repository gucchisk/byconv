package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/gucchisk/bytestring"
	"github.com/gucchisk/byconv/flags"
)

func main() {
	var inEnc, outEnc, filename string
	app := &cli.App{
		Name: "byconv",
		Usage: "byte converter (ascii, hex, base64)",
		Description: "convert byte string of file or stdin",
		UsageText: "byconv [options]",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "input",
				Usage: "input encoding",
				Aliases: []string{"i"},
				Destination: &inEnc,
			},
			&cli.StringFlag{
				Name: "output",
				Usage: "output encoding",
				Aliases: []string{"o"},
				Destination: &outEnc,
			},
			&cli.StringFlag{
				Name: "file",
				Usage: "input file",
				Aliases: []string{"f"},
				Destination: &filename,
			},
		},
		Action: func(c *cli.Context) error {
			if inEnc != "" {				
				fmt.Printf("input: %s\n", inEnc)
			}

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

			in, err := flags.NewEncoding(inEnc)
			if err != nil {
				return err
			}
			bytes, err := bytestring.NewBytes(byteArray, bytestring.SetEncoding(bytestring.Encoding(in)))
			if err != nil {
				return err
			}
			out, err := flags.NewEncoding(outEnc)
			if err != nil {
				return err
			}

			var result string
			switch out {
			case bytestring.Ascii:
				result = bytes.String()
			case bytestring.Hex:
				result = bytes.HexString()
			case bytestring.Base64:
				result = bytes.Base64()
			case bytestring.Base64URL:
				result = bytes.Base64URL()
			}
			fmt.Printf("%s", result)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		os.Exit(1)
	}
}
