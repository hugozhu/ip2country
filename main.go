package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"github.com/oschwald/maxminddb-golang"
)


func main() {
	db, err := maxminddb.Open("Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    var ipAddress string
    if len(os.Args) > 1 {
        ipAddress = os.Args[1]
    } else {
        fmt.Scanln(&ipAddress)
    }

	ip := net.ParseIP(ipAddress)

	record := make(map[string]string)
	if err := db.Lookup(ip, &record); err != nil || len(record) == 0 {
		log.Fatal(err)
	}
	fmt.Printf("%s %s\n", ip, record["country"])
}