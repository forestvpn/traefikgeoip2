package traefikgeoip2

import (
	"fmt"
	"net"
	"strconv"

	"github.com/IncSW/geoip2"
)

// Unknown constant for undefined data.
const Unknown = "XX"

// DefaultDBPath default GeoIP2 database path.
const DefaultDBPath = "GeoLite2-Country.mmdb"

// DefaultIPHeader default ip header name.
const DefaultIPHeader = ""

const (
	// CountryHeader country header name.
	CountryHeader = "X-GeoIP2-Country"
	// RegionHeader region header name.
	RegionHeader = "X-GeoIP2-Region"
	// CityHeader city header name.
	CityHeader = "X-GeoIP2-City"
	// LatitudeHeader latitude header name.
	LatitudeHeader = "X-GeoIP2-Latitude"
	// LongitudeHeader longitude header name.
	LongitudeHeader = "X-GeoIP2-Longitude"
)

// GeoIPResult GeoIPResult.
type GeoIPResult struct {
	country   string
	region    string
	city      string
	latitude  string
	longitude string
}

// LookupGeoIP2 LookupGeoIP2.
type LookupGeoIP2 func(ip net.IP) (*GeoIPResult, error)

// CreateCityDBLookup CreateCityDBLookup.
func CreateCityDBLookup(rdr *geoip2.CityReader) LookupGeoIP2 {
	return func(ip net.IP) (*GeoIPResult, error) {
		rec, err := rdr.Lookup(ip)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		retval := GeoIPResult{
			country:   rec.Country.ISOCode,
			region:    Unknown,
			city:      Unknown,
			latitude:  Unknown,
			longitude: Unknown,
		}
		if city, ok := rec.City.Names["en"]; ok {
			retval.city = city
		}
		if rec.Subdivisions != nil {
			retval.region = rec.Subdivisions[0].ISOCode
		}
		if rec.Location.Latitude != 0 {
			retval.latitude = strconv.FormatFloat(rec.Location.Latitude, 'f', -1, 64)
		}
		if rec.Location.Longitude != 0 {
			retval.longitude = strconv.FormatFloat(rec.Location.Longitude, 'f', -1, 64)
		}

		return &retval, nil
	}
}

// CreateCountryDBLookup CreateCountryDBLookup.
func CreateCountryDBLookup(rdr *geoip2.CountryReader) LookupGeoIP2 {
	return func(ip net.IP) (*GeoIPResult, error) {
		rec, err := rdr.Lookup(ip)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		retval := GeoIPResult{
			country:   rec.Country.ISOCode,
			region:    Unknown,
			city:      Unknown,
			latitude:  Unknown,
			longitude: Unknown,
		}
		return &retval, nil
	}
}
