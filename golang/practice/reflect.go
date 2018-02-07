package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

func main() {
	var num float64 = 1.2345
	numType := reflect.TypeOf(num)
	numValue := reflect.ValueOf(num)
	numPtr := reflect.ValueOf(&num)

	fmt.Println("type: ", numType)
	fmt.Println("value: ", numValue)
	fmt.Println("ptr: ", numPtr)

	//强制类型转换
	convertPointer := numPtr.Interface().(*float64)
	convertValue := numValue.Interface().(float64)
	fmt.Println(convertPointer)
	fmt.Println(convertValue)

	newValue := numPtr.Elem() //只有指针才能赋值
	fmt.Println("type of numptr:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet()) //是否可以设置
	//重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of num:", num)

	fmt.Println("----------------------------------------------")
	u := User{1, "Arvin", 25}

	DoFiledAndMethod(u)
}

func DoFiledAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is : ", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get Value is : ", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

	fmt.Println("-----------------------------------------------")
	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)} //参数类型数量，类型保持一致
	methodValue.Call(args)                                                   //调用方法

	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)
}
