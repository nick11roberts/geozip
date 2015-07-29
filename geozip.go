//Package geozip (some description here)
package geozip

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

//Encode (some description here)
func Encode(latitude, longitude float64, validate bool, precision int) int64 {
	if !Validate(latitude, longitude) {
		return 0
	}
	latitudeShifted := decimal.NewFromFloat(latitude).Add(decimal.NewFromFloat(90.0))
	longitudeShifted := decimal.NewFromFloat(longitude).Add(decimal.NewFromFloat(180.0))
	latString := latitudeShifted.String()
	lonString := longitudeShifted.String()
	latString = strings.Replace(latString, ".", "", 1)
	lonString = strings.Replace(lonString, ".", "", 1)
	bucketString := zip(latString, lonString)
	bucket, err := strconv.ParseInt(bucketString, 10, 64)
	if err != nil {
		fmt.Errorf("Error parsing zipped string to int64")
	}
	//Fix precision here
	return bucket
}

//Decode (some description here)
func Decode(bucket int64) (float64, float64) {
	var latitude, longitude float64
	return latitude, longitude
}

//Validate (some description here)
func Validate(latitude, longitude float64) bool {
	if (latitude < 90.0 || latitude > -90.0) && (longitude < 180.0 || longitude > 180.0) {
		return true
	}
	return false

}

func zip(latDigits, lonDigits string) string {
	var bucketDigits string
	latDigits = resize(latDigits)
	lonDigits = resize(lonDigits)
	for i := 0; i < 9; i++ {
		bucketDigits += string(latDigits[i])
		bucketDigits += string(lonDigits[i])
	}
	return bucketDigits
}

func resize(component string) string {
	for len(component) > 9 {
		component = component[0 : len(component)-1]
	}
	for len(component) < 9 {
		component = "0" + component
	}
	return component
}
