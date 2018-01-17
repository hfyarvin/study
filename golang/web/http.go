package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	// "os"
	// "io"
)

func get(url string, headers map[string]string) *http.Response {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return resp
}
func main() { //生成client 参数为默认
	AddUser()
}

// body, err := ioutil.ReadAll(response.Body)
// if err != nil {
// 	panic(err)
// }
// var ret interface{}
// err = json.Unmarshal(body, &ret)

func testGet() {
	url := "http://localhost:9000/******"
	h := make(map[string]string)
	h["Accept"] = "application/json"
	h["Authorization"] = ""

	response := get(url, h)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var ret interface{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", ret)
}

func post(url string, headers map[string]string, bodyStr string) *http.Response {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr,
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(bodyStr))

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return resp
}

type UserContact struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:""email`
	Country   string `json:"country"`
}

func postTest() {
	url := "http://localhost:9000/******"
	h := make(map[string]string)
	h["Accept"] = "application/json"
	h["Content-Type"] = "application/json"
	h["Authorization"] = ""
	//
	b := UserContact{
		FirstName: "****",
		LastName:  "****",
		Email:     "****98207@qq.com",
		Country:   "Cn",
	}

	response := post(url, h, JsonToStr(b))
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var ret interface{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", ret)
}

func JsonToStr(b interface{}) string {
	bs, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(bs)
}
