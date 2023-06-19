package timezonex

type Zone struct {
	CountryCode string
	Name        string
}

// Country define a country include code,name,and timezone list
type Country struct {
	Code string
	Name string
	//contain time zone list
	Zones []Zone
}
