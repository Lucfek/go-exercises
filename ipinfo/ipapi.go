package ipapi

import (
    "net/http"
    "net"
    "encoding/json"
    "log"
    "errors"
    "io/ioutil"
)


type IpApiFields struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Country string `json:"country"`
    CountryCode string `json:"countryCode"`
    Region string `json:"region"`
    RegionName string `json:"regionName"`
    City string `json:"city"`
    District string `json:"district"`
    Zip string `json:"zip"`
    Lat float64 `json:"lat"`
    Lon float64 `json:"lon"`
    Timezone string `json:"timezone"`
    Isp string `json:"isp"`
    As string `json:"as"`
    Reverse string `json:"reverse"`
    Mobile bool `json:"mobile"`
    Proxy bool `json:"proxy"`
    Query string `json:"query"`
    Org string `json:"org"`
}


//Returns all Info(struct) about given IP address
func GetIpInfo(addr string) IpApiFields  {
    
    if net.ParseIP(addr) == nil {
        err := errors.New("Wrong address syntax")
        log.Fatal(err)
    }
    
    url := "http://ip-api.com/json/" + addr
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
    
    fields := IpApiFields{}
    
    err = json.Unmarshal([]byte(body), &fields)
	if err != nil {
		log.Fatal(err)
	}
    
    if fields.Message != "" {
        err = errors.New(fields.Message)
        log.Fatal(err)
    }
    
    return fields  
}