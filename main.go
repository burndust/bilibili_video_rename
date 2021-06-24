package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type InfoStruct struct{
	PartName string
}

func main() {
	for{
		fmt.Println("请输入文件路径：")
		var root string
		fmt.Scanf("%s\n",&root)
		rename(root)
	}
}

func rename(root string){
	infoMap := make(map[int]string, 20)
	videoMap := make(map[int]string, 20)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		relativePath := strings.TrimPrefix(path, root)
		splitRP := strings.Split(relativePath, "\\")
		if strings.HasSuffix(relativePath, ".info") {
			index,_ := strconv.Atoi(splitRP[1])
			infoMap[index] = path
		} else if strings.HasSuffix(relativePath, ".mp4") {
			index,_ := strconv.Atoi(splitRP[1])
			videoMap[index] = path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for index, file := range infoMap {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}

		infoStruct := &InfoStruct{}
		err = json.Unmarshal(content, infoStruct)
		if err != nil {
			fmt.Println("json unmarshal failed!")
			return
		}
		paths, _ := filepath.Split(file)
		newFile := paths + infoStruct.PartName + ".mp4"
		oserr := os.Rename(videoMap[index],newFile)
		if oserr != nil {
			fmt.Println("reanme file failed, err:", videoMap[index],newFile,oserr)
			panic(err)
		}
	}
	fmt.Println("执行完成")
}
