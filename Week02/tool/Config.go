package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/**
 * 配置项
 */
type Config struct{
	Database DatabaseConig `json:"database"`
}


type DatabaseConig struct {
	Driver string `json:"driver"`
	User string `json:"user"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
	DbName string `json:"db_name"`
	Charset string `json:"charset"`
	ShowSql bool `json:"show_sql"`
}


var _cfg *Config = nil

func GetConfig() *Config{
	return _cfg
}

/**
 * json转化
 */
func ParseConfig(path string)(*Config,error){
	fmt.Println(path)
	file,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err :=decoder.Decode(&_cfg);err != nil{
		return nil,err
	}
	return _cfg,nil
}