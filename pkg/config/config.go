// Package config implements the configuration parsers used for the authentication

package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/go-yaml/yaml"
)

const (
	defaultConfig = `port: 7777
db:
  type: "sqlite3"
  addr: ""
  name: "./test.db"
  user: ""
  pass: ""
`
)

// Config is the structure containing all information needed to configure and run the server.
type DefaultConfig struct {
	Port     int `yaml:"port"`
	Database struct {
		Kind         string `yaml:"type" mapstructure:"type"` // Database type (i.e. sqlite3, postgres, mysql, etc.)
		Addr         string `yaml:"addr" mapstructure:"addr"` // Address of database. If database type is "sqlite3", then leave this field blank
		DatabaseName string `yaml:"name" mapstructure:"name"` // Database name, or in the case of sqlite3, the path to the database
		Username     string `yaml:"user" mapstructure:"user"` // Database username
		Password     string `yaml:"pass" mapstructure:"pass"` // Database password
	} `yaml:"db" mapstructure:"db"`
}

func GetConfig(path, filename string) (*DefaultConfig, error) {
	if !tryConfig(path, filename) {
		log.Printf("Failed to load config file, attempting to generate file.\n")
	}
	if err := generateDefaultConfig(path, filename, true); err != nil {
		log.Printf("failed to load config file, resorting to default\n")

		c := DefaultConfig{}
		if err := yaml.Unmarshal([]byte(defaultConfig), &c); err != nil {
			panic(err)
		}
	}

	v, err := readConfig(path, filename, map[string]interface{}{
		"port": 7777,
		"database": map[string]interface{}{
			"type": "sqlite3",
			"addr": "",
			"name": "./test.db",
			"user": "",
			"pass": "",
		},
	})
	if err != nil {
		panic(err)
	}

	c := DefaultConfig{}
	err = v.Unmarshal(&c)
	return &c, err

}

func tryConfig(path, filename string) bool {
	if _, err := os.Stat(path + "/" + filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func generateDefaultConfig(path, filename string, force bool) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if force {
			os.Mkdir(path, 0755)
		} else {
			return err
		}
	}

	err := ioutil.WriteFile(path+"/"+filename, []byte(defaultConfig), 0644)
	return err
}

func readConfig(path, filename string, defaults map[string]interface{}) (*viper.Viper, error) {

	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigFile(filename)
	v.AddConfigPath(path)
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}
