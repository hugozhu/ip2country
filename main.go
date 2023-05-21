package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"github.com/oschwald/geoip2-golang"
)


func main() {
	db, err := geoip2.Open("Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	args := os.Args[1:]
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(args[0])
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", record.Country.IsoCode)
}