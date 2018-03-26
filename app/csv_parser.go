package app

import (
	"github.com/Lacka90/gophulate/interfaces"
	"github.com/gocarina/gocsv"
)

// ParseCSV - parse csv data
func ParseCSV(data string) []*interfaces.Record {
	clients := []*interfaces.Record{}

	if err := gocsv.UnmarshalString(data, &clients); err != nil {
		panic(err)
	}

	return clients
}
