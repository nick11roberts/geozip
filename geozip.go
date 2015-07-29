//Package geozip (some description here)
package geozip

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

const maxPrecision int = 18

//Encode (some description here)
func Encode(latitude, longitude float64, validate bool, precision int) int64 {
	if !Validate(latitude, longitude) {
		return 0
	}
	latitudeShifted := decimal.NewFromFloat(latitude).Add(decimal.NewFromFloat(90.0))
	longitudeShifted := decimal.NewFromFloat(longitude).Add(decimal.NewFromFloat(180.0))
	latString := latitudeShifted.String() + ".0"
	lonString := longitudeShifted.String() + ".0"
	latParts := strings.Split(latString, ".")
	lonParts := strings.Split(lonString, ".")
	latString = resizeCharacteristic(latParts[0]) + resizeMantissa(latParts[1])
	lonString = resizeCharacteristic(lonParts[0]) + resizeMantissa(lonParts[1])
	bucketString := zip(latString, lonString)
	bucket, err := strconv.ParseInt(bucketString, 10, 64)
	if err != nil {
		fmt.Errorf("Error parsing zipped string to int64")
	}
	for i := 0; i < maxPrecision-precision; i++ {
		bucket /= 10
	}
	for i := 0; i < maxPrecision-precision; i++ {
		bucket *= 10
	}
	return bucket
}

//Decode (some description here)
func Decode(bucket int64) (float64, float64) {
	var latitude, longitude float64
	return latitude, longitude
}

//Validate (some description here)
func Validate(latitude, longitude float64) bool {
	if latitude < 90.0 && latitude > -90.0 && longitude < 180.0 && longitude > -180.0 {
		return true
	}
	return false

}

func zip(latDigits, lonDigits string) string {
	var bucketDigits string
	for i := 0; i < 9; i++ {
		bucketDigits += string(latDigits[i])
		bucketDigits += string(lonDigits[i])
	}
	return bucketDigits
}

func resizeCharacteristic(characteristic string) string {
	for len(characteristic) < 3 {
		characteristic = "0" + characteristic
	}
	return characteristic
}

func resizeMantissa(mantissa string) string {
	for len(mantissa) > 6 {
		mantissa = mantissa[0 : len(mantissa)-1]
	}
	for len(mantissa) < 6 {
		mantissa = mantissa + "0"
	}
	return mantissa
}
