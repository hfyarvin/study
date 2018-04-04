package main

import (
	"github.com/tiaguinho/gosoap"
	"fmt"
)

type GetGeoIPResponse struct {
	GetGeoIPResult GetGeoIPResult
}

type GetGeoIPResult struct {
	ReturnCode        string
	IP                string
	ReturnCodeDetails string
	CountryName       string
	CountryCode       string
}

var (
	r GetGeoIPResponse
)

func main() {
	soap, err := gosoap.SoapClient("http://www.webservicex.net/geoipservice.asmx?WSDL")
	if err != nil {
		fmt.Errorf("error not expected: %s", err)
	}

	params := gosoap.Params{
		"IPAddress": "8.8.8.8",
	}

	err = soap.Call("GetGeoIP", params)
	if err != nil {
		fmt.Errorf("error in soap call: %s", err)
	}

	soap.Unmarshal(&r)
	if r.GetGeoIPResult.CountryCode != "USA" {
		fmt.Errorf("error: %+v", r)
	}

	fmt.Println(r.GetGeoIPResult.CountryCode)
	fmt.Println(r.GetGeoIPResult.ReturnCodeDetails)
}