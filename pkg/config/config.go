package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

// Config model
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	DB struct {
		User        string `yaml:"user"`
		Pass        string `yaml:"pass"`
		Name        string `yaml:"name"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		Mode        string `yaml:"mode"`
		AutoMigrate bool   `yaml:"auto_migrate"`
	} `yaml:"db"`
	Minio struct {
		Scheme    string `yaml:"scheme"`
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
		UseSSL    bool   `yaml:"useSSL"`
		Bucket  string  `yaml:"bucket"`
	} `yaml:"minio"`
}

// Get - Config initializer
func Get() *Config {
	f, err := os.Open("configs/config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return &config
}
