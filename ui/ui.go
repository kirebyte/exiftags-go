package ui

import (
	"github.com/urfave/cli/v2"
)

var allowRecursiveScan bool
var format string

func New(actionfunc cli.ActionFunc) *cli.App {
	return &cli.App{
		Name:        "exiftags-go",
		Description: "exiftags-go is a command-line utility for extracting and displaying EXIF metadata tags from image files. It supports a variety of image formats and allows users to specify the output format for the extracted tags. This tool is useful for photographers, digital forensics investigators, and anyone else who needs to view the metadata embedded in an image file.",
		Usage:       "Extract and display EXIF tags from an image",
		UsageText:   "exiftags-go [options] file...",
		Action:      actionfunc,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "recursive",
				Value:       false,
				Usage:       "if any of the file parameters is a directory then crawl inside looking for more images",
				Aliases:     []string{"r"},
				Destination: &allowRecursiveScan,
			},
			&cli.StringFlag{
				Name:        "format",
				Value:       "raw",
				Usage:       "set output format to: [csv|html|json|raw]. Example: --format csv",
				Aliases:     []string{"f"},
				Destination: &format,
			},
		},
		Authors: []*cli.Author{
			{
				Name:  "Erick Gustavo López Siordia (KireByte)",
				Email: "kirebyte@gmail.com",
			},
		},
		Copyright: "Copyright © 2023 KireByte MIT License: https://www.mit.edu/~amini/LICENSE.md \nThis is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent permitted by law.",
	}
}
