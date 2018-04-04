package main

import (
	"fmt"
	"github.com/tiaguinho/gosoap"
)

type GetAllDataResponse struct {
	GetAllDataResult GetAllDataResult
}

type GetAllDataResult struct {
	DateTime    string `xml:"data_time"`
	GatewayLog  string `xml:"gateway_log"`
	SensorName  string `xml:"sensor_name"`
	ChannelName string `xml:"channel_name"`
	Value       string `xml:"value"`
}

var (
	r GetAllDataResponse
)

func main() {
	soap, err := gosoap.SoapClient("http://iot.klha.net:8080/services/iotDataService?wsdl")
	if err != nil {
		fmt.Errorf("error not expected: %s", err)
	}

	params := gosoap.Params{
		"username": "nbdemo",
		"password": "123456",
	}

	err = soap.Call("CurAllData", params)
	if err != nil {
		fmt.Errorf("error in soap call: %s", err)
	}

	err = soap.Unmarshal(&r)
	// if r.CurAllDataResult.CountryCode != "USA" {
	// 	fmt.Errorf("error: %+v", r)
	// }
	if err != nil {
		fmt.Errorf("error in soap call: %s", err)
	}
	fmt.Println("=====")
	fmt.Println(r.GetAllDataResult.Value)
}
