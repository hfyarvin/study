package rsa

import (
	"crypto"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
	"toast/rsa"
)

var cipher rsa.Cipher

func init() {
	client, err := rsa.NewDefault(`-----BEGIN PRIVATE KEY-----
私钥信息
-----END PRIVATE KEY-----`, `-----BEGIN PUBLIC KEY-----
公钥信息
-----END PUBLIC KEY-----`)

	if err != nil {
		fmt.Println(err)
	}

	cipher = client
}

func Test_DefaultClient(t *testing.T) {

	cp, err := cipher.Encrypt([]byte("测试加密解密"))
	if err != nil {
		t.Error(err)
	}
	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp, err := cipher.Decrypt(ppBy)

	fmt.Println(string(pp))
}

func Test_Sign_DefaultClient(t *testing.T) {

	src := "测试签名验签"

	signBytes, err := cipher.Sign([]byte(src), crypto.SHA256)
	if err != nil {
		t.Error(err)
	}
	sign := hex.EncodeToString(signBytes)
	fmt.Println(sign)

	signB, err := hex.DecodeString(sign)

	errV := cipher.Verify([]byte(src), signB, crypto.SHA256)
	if errV != nil {
		t.Error(errV)
	}
	fmt.Println("verify success")
}
