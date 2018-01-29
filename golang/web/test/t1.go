package main

import (
	"fmt"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
func web1() {
	http.HandleFunc("/", web1_handler)
	err := http.ListenAndServe(":9999", nil)
	checkErr(err)
}

func web1_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ...")
}

func web2() {
	http.HandleFunc("/", web2_handler)
	err := http.ListenAndServe(":9999", nil)
	checkErr(err)
}

func web2_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("url_long:", r.Form["url_long"]) //参数url_long为query参数
	for k, v := range r.Form {
		fmt.Println(fmt.Sprintf("k:%v,v:%v", k, v))
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
}
