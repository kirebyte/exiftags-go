package controller

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kirebyte/exiftags-go/entities"
	"github.com/kirebyte/exiftags-go/exif"
	"github.com/urfave/cli/v2"
)

// D.O. controller ;)
// This module is basically a PoC demo of this MVC inspired architecture.
// I want to make a complete solution later and all I will need to to is replace this controller
// and expand the EXIF package that works as the model.

func NewExifActionFunc() cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		// stop if there's not at least 1 file
		if c.Args().Len() == 0 {
			return cli.Exit("exiftags-go: No filenames were provided. At least 1 expected", 2)
		}

		filesList := c.Args().Slice()
		// on recursive flag walk inside directories and update the files list
		if c.Bool("recursive") {
			filesList, err = listNestedFiles(filesList)
			if err != nil {
				return err
			}
		}

		var records []entities.ExifRecord
		// get the required information, in this case: filename, latitude and longitude
		for _, file := range filesList {
			if gpsTag, err := exif.GetGpsInfo(file); err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", file, err.Error())
			} else {
				records = append(records, entities.ExifRecord{Filename: file, Latitude: gpsTag.Latitude.String(), Longitude: gpsTag.Longitude.String()})
			}
		}

		// print
		switch c.String("format") {
		case "raw": //default hardcoded case
			entities.PrintRaw(records)
		case "csv":
			entities.PrintCsv(records)
		case "html":
			entities.PrintHtml(records)
		case "json":
			entities.PrintJson(records)
		}

		return
	}
}

func listNestedFiles(paths []string) ([]string, error) {
	var files []string
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}
