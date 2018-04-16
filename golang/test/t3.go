package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type sensorData struct {
	Value          float64
	Electric       float64
	SignalStrength float64
}

func (s *sensorData) UnmarshalBinary(b []byte) error {
	if len(b) != 12 {
		return errors.Errorf("data length must be 12,expected %d", len(b))
	}
	strval := fmt.Sprintf("%s", b[:5])
	//fmt.Println("test:", strval)
	floatVal, err := strconv.ParseFloat(strval, 64)
	if err != nil {
		return err
	}
	s.Value = floatVal

	strElectric := fmt.Sprintf("%s", b[5:8])
	//fmt.Println("test:", strElectric)
	floatElectric, err := strconv.ParseFloat(strElectric, 64)
	if err != nil {
		return err
	}
	s.Electric = floatElectric

	strSignal := fmt.Sprintf("%s", b[8:])
	//fmt.Println("test", strSignal)
	floatSignal, err := strconv.ParseFloat(strSignal, 64)
	if err != nil {
		return err
	}
	s.SignalStrength = floatSignal

	return nil
}

func main()  {
	s := new(sensorData)
err :=s.UnmarshalBinary([]byte("3031322E3334"))
if err !=nil{
	fmt.Println(err.Error())
}
fmt.Println(s)

// //-----------------------
// 	strval := fmt.Sprintf("%s", []byte("31322E3334"))
// 	//fmt.Println("test:", strval)
// 	floatVal, err := strconv.ParseFloat(strval, 64)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(floatVal)
}