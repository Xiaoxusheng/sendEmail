package main

import (
	"encoding/json"
	"fmt"
	"github.com/Xiaoxusheng/sendEmail/utils"
	"log"
	"os"
)

func main() {
	//  解析config.json
	config := new(utils.Config)
	//	读取
	file, err := os.Open("./config.json")
	if err != nil {
		log.Println("配置文件读取错误：" + err.Error())
		return
	}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		log.Println("json解析失败" + err.Error())
		return
	}
	fmt.Println(config)
	utils.Loop(config)
}
