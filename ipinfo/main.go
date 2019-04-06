package main

import (
    "fmt"
    "flag"
    "ipinfo/ipapi"
)



func main() {
    ipPtr  := flag.String("ip", "", "An IP address you want to check")
    geoPtr := flag.Bool("geo", false, "")
    flag.Parse()

        
    resp := ipapi.GetIpInfo(*ipPtr)
    if *geoPtr {
        fmt.Printf("Loc: %.4f, %.4f", resp.Lat, resp.Lon)
    } else {
        fmt.Printf("IP Address: %s\nOragnization: %s\nCity: %s\nRegion: %s\nCountry: %s\nLoc: %.4f, %.4f\nPostal: %s", resp.Query, resp.Org, resp.City, resp.RegionName, resp.Country, resp.Lat, resp.Lon, resp.Zip)
    }
}