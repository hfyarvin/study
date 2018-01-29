package main

import (
	// "./tpl" //调用同一文件夹下的目录
	// "strings"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	} else if r.URL.Path == "/login" {
		login(w, r)
		return
	} else if r.URL.Path == "/selectoption" {
		selectOption(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello myroute")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./tpl/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func selectOption(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./tpl/selectOption.gtpl")
		log.Panicln(t.Execute(w, nil))
		// 验证输入
		slice := []string{"apple", "pear", "banana"}
		v := r.Form.Get("fruit")
		for _, item := range slice {
			if item == v {
				fmt.Println("the param is : ", v)
			}
		}
	} else {
		// 请求的是登陆数据，那么执行登陆的逻辑判断
		// fmt.Println("username:", r.Form["username"])
		// fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
