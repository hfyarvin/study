//Golang中encoding的用法
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"os"
)

func goxml() {
	type Address struct {
		City, State string
	}

	//xml元素节点增加属性，需要重新定制结构体
	type ElementWithAttr struct {
		AttrOne string `xml:"attrone,attr"`
		AttrTwo string `xml:"attrtow,attr"`
	}

	type Person struct {
		XMLName   xml.Name        `xml:"person"`
		Id        int             `xml:"id,attr"`
		Sex       string          `xml:"sex,attr"`
		FirstName string          `xml:"name>first"`
		LastName  string          `xml:"name>last"`
		Age       int             `xml:"age"`
		Height    float32         `xml:"height,omitempty"`
		EWA       ElementWithAttr `xml:elementwithattr`
		//这里写子节点或者属性的时候（凡是不是一个单独字符串），必须使用""来进行包封，否则反射的时候认不出来
		EWAChild string `xml:"elementwithattr>ewachild"`
		Married  bool
		Comment  string `xml:",comment"`
		Address
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42, Sex: "Female"}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}
	v.EWA = ElementWithAttr{"AttribteOne", "AttributeTwo"}
	v.EWAChild = "ElementWithAttrChildNode"

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func main() {
	pi := 3.1415926
	buf := bytes.Buffer{}
	//使用的是小编码，低地址对应低字节
	binary.Write(&buf, binary.LittleEndian, &pi)
	//常量浮点数默认是float64
	fmt.Printf("buf=%#v\n", buf.Bytes())

	var rpi float64
	binary.Read(&buf, binary.LittleEndian, &rpi)
	fmt.Printf("rpi=%#v\n", rpi)

	//和python的binascii库的作用一样，16进制和ascii字符之间的转换
	src := "Go is good language!我们都一样"
	//根据编码后长度来分配缓存空间
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, []byte(src))
	fmt.Printf("dst=%v\n", dst)
	fmt.Printf("dst(hex)=%s\n", string(dst))

	//根据解码长度来分配缓存空间
	dst2 := make([]byte, hex.DecodedLen(len(dst)))
	hex.Decode(dst2, dst)
	fmt.Printf("dst2=%v\n", dst2)
	fmt.Printf("dst2.(string)=%s\n", string(dst2))

	goxml()
}
