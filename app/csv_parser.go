package app

import "github.com/gocarina/gocsv"

// ParseCSV - parse csv data
func ParseCSV(data string) []*Record {
	clients := []*Record{}

	if err := gocsv.UnmarshalString(data, &clients); err != nil {
		panic(err)
	}

	return clients
}
