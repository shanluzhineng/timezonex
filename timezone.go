package timezonex

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

// random a timezone for country
func RandomTimeZone(countryCode string) *Zone {
	if len(countryCode) <= 0 {
		return nil
	}
	countryCode = strings.ToUpper(countryCode)
	country, found := GetCountry(countryCode)
	if !found || len(country.Zones) <= 0 {
		return nil
	}
	if len(country.Zones) == 1 {
		return &country.Zones[0]
	}
	len := len(country.Zones)
	return &country.Zones[randInt32(0, len)]
}

func ZoneList() []string {
	return zoneList
}

// get country by code
func ZoneListByCountry(code string) *Country {
	if code == "" {
		return nil
	}
	code = strings.ToUpper(code)
	country, ok := mapped[code]
	if ok {
		return &country
	}
	return nil
}

func ZoneIsExist(zone string) bool {
	_, ok := zoneMap[zone]
	return ok
}

// 在2个值之间随机一个值
func randInt32(minValue int, maxValue int) int {
	rand.Seed(time.Now().UnixNano())
	randValue := rand.Intn(maxValue)
	if minValue < 0 {
		return randValue - int(math.Abs(float64(minValue)))
	}
	return randValue
}
