package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//read folder
	folder := "E:/work/go/golang_learn_detail/golang_learn_detail"
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		//get file name
		fmt.Println(file.Name())
	}
	_, err, info := PathExist("E:/work/go/golang_learn_detail/golang_learn_detail/t.go")
	if err == nil {
		fmt.Println(info)
	}
	//移除文件或者文件夹
	removeErr := os.RemoveAll("E:/work/go/golang_learn_detail/golang_learn_detail/t.go")
	if removeErr != nil {
		fmt.Println("RemoveDstPathFolder  err: ", removeErr)
	}

	emptyList, findErr := GetALlEmptyFolder("E:/work/go/golang_learn_detail/golang_learn_detail")
	if findErr == nil {
		for _, item := range emptyList {
			fmt.Println(item)
		}
	} else {
		fmt.Println(findErr)
	}
}

//判断文件夹是否存在
func PathExist(path string) (bool, error, os.FileInfo) {
	info, err := os.Stat(path)
	// 为空则表示文件或者文件夹存在
	if err == nil {
		return true, nil, info
	}
	// 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	if os.IsNotExist(err) {
		return false, nil, info
	}
	return false, err, info
}

//找到某个文件夹下面的所有空文件夹
func GetALlEmptyFolder(folder string) ([]string, error) {
	fileNames := make([]string, 0)
	dirNames := make([]string, 0)
	var emptyFolderList []string
	walkDir := folder
	//遍历文件夹并把文件或文件夹名称加入相应的slice
	err := filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirNames = append(dirNames, path)
		} else {
			fileNames = append(fileNames, path)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	//把所有文件名称连接成一个字符串
	fileNamesAll := strings.Join(fileNames, "")
	for i := len(dirNames) - 1; i >= 0; i-- {
		//文件夹名称不存在文件名称字符串内说明是个空文件夹
		if !strings.Contains(fileNamesAll, dirNames[i]) {
			// fmt.Printf("%s is empty\n", dirNames[i])
			emptyFolderList = append(emptyFolderList, dirNames[i])
			/*
			   err:=os.Remove(dirNames[i])
			   if err!=nil{
			       fmt.Println(err)
			   }*/
		}
	}
	return emptyFolderList, nil
}
