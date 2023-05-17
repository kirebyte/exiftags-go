package exif

import (
	exif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/kirebyte/exiftags-go/entities"
)

func GetGpsInfo(filepath string) (gpsInfo *entities.GpsInfo, err error) {
	rawExif, err := exif.SearchFileAndExtractExif(filepath)
	if err != nil {
		return nil, err
	}

	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, err
	}

	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(im, ti, rawExif)
	if err != nil {
		return nil, err
	}

	rootIfd := index.RootIfd

	gpsIfd, err := rootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return nil, err
	}

	initialGi, err := gpsIfd.GpsInfo()
	if err != nil {
		return nil, err
	}

	gpsLocation := &entities.GpsInfo{
		Latitude: entities.GpsLocation{
			Degrees:     int(initialGi.Latitude.Degrees),
			Minutes:     int(initialGi.Latitude.Minutes),
			Seconds:     initialGi.Latitude.Seconds,
			Orientation: initialGi.Latitude.Orientation,
		},
		Longitude: entities.GpsLocation{
			Degrees:     int(initialGi.Longitude.Degrees),
			Minutes:     int(initialGi.Longitude.Minutes),
			Seconds:     initialGi.Longitude.Seconds,
			Orientation: initialGi.Longitude.Orientation,
		},
	}

	return gpsLocation, nil
}
