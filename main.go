package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"github.com/oschwald/geoip2-golang"
	"github.com/ipinfo/mmdbctl"
    "net/http"
    "io/ioutil"
    "encoding/json"	
)


func main() {
	db, err := mmdbctl.Open("Country.mmdb")
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
	record, err := db.City(ip)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %s\n", ip, record.Country.IsoCode)
}

type IpInfo struct {
    Country string `json:"country"`
}

func ip_info(ip string) string {
    url := "https://ipinfo.io/"+ip+"?token=xxxxx"
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    var ipInfo IpInfo
    err = json.Unmarshal(body, &ipInfo)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    return ipInfo.Country
}