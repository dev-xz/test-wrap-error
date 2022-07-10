package conf

import (
	"encoding/json"
	"os"
)

type DBConf struct {
	Type     string `json:"type,omitempty"`
	Host     string `json:"host,omitempty" json:"host,omitempty"`
	Port     int    `json:"port,omitempty" json:"port,omitempty"`
	Username string `json:"username,omitempty" json:"username,omitempty"`
	Passwd   string `json:"passwd,omitempty" json:"passwd,omitempty"`
	DbName   string `json:"db_name,omitempty" json:"db_name,omitempty"`
}

var DBConfs map[string]DBConf

func init() {
	// 打开文件
	file, _ := os.Open("conf/db.json")
	// 关闭文件
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&DBConfs)
	if err != nil {
		panic(err)
	}
}
