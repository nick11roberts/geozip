package geozip

import "testing"

var geozipTests = []struct {
	latitude  float64
	longitude float64
	validate  bool
	precision int
	bucket    int64
}{
	{-34.783467, 128.294109, true, 18, 35058221964513039},
}

func TestEncode(t *testing.T) {
	for _, tableEntry := range geozipTests {
		bucket := Encode(tableEntry.latitude, tableEntry.longitude, tableEntry.validate, tableEntry.precision)
		if bucket != tableEntry.bucket {
			t.Errorf("Encode(%v, %v, %v, %v) = %v, want %v", tableEntry.latitude, tableEntry.longitude, tableEntry.validate, tableEntry.precision, bucket, tableEntry.bucket)
		}
	}
}

func BenchmarkGeozip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(-34.783467, 128.294109, true, 18)
	}
}
