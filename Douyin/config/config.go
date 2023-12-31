package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Resource Resource `yaml:"resource"`
}

type Mysql struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Ipaddress string `yaml:"ipaddress"`
	Port      string `yaml:"port"`
	Dbname    string `yaml:"dbname"`
}

type Redis struct {
	Ipaddress string `yaml:"ipaddress"`
	Port      string `yaml:"port"`
	Maxidle   int    `yaml:"maxidle"`
	Maxactive int    `yaml:"maxactive"`
}

type Resource struct {
	Ipaddress string `yaml:"ipaddress"`
	Port      string `yaml:"port"`
}

var C Config

// ConfInit 将读取的yaml文件解析为响应的 struct
func ConfInit() error {
	yamlFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = yaml.Unmarshal(yamlFile, &C)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
