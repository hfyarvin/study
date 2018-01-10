/**
 * 网易: smtp.163.com(ssl:465/994,非ssl: 25)
 * Gmail: stmp.gmail.com(ssl:465)
 */
package main

import (
	"fmt"

	"github.com/go-gomail/gomail"
)

func check(e error) {
	if e != nil {
		panic(e) //panic 和recover
	}
}

func main() {
	send3()
}

func send3() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "sender@gmail.com", "arvin")
	m.SetHeader("To", m.FormatAddress("receiver@qq.com", "arvin"))
	m.SetHeader("Cc", m.FormatAddress("sender@gmail.com", "arvin"))
	m.SetHeader("Subject", "这是一个文件主题")
	m.SetBody("text/html", "正文，你好，这是正文，这是正文。")
	d := gomail.NewPlainDialer("smtp.gmail.com", 465, "sender@gmail.com", "*****")
	// d := gomail.NewPlainDialer("smtp.163.com", 25, "17688138833@163.com", "arvin171017")
	err := d.DialAndSend(m)
	check(err)
	fmt.Printf("successful")
}
