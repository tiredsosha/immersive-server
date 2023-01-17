package tools

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

type conf struct {
	OscIp    string `yaml:"oscIp"`
	OscPort  string `yaml:"oscPort"`
	HttpPort string `yaml:"httpPort"`
}

func getConf(file string, cnf interface{}) error {
	yamlFile, err := os.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, cnf)
	}

	return err
}

func validateConf(cfg *conf) error {
	var err error
	v := reflect.ValueOf(*cfg)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := typeOfS.Field(i).Name
		value := v.Field(i).Interface()
		if value == "" {
			err = errors.New("config: " + strings.ToLower(field) + " field is emtpy/nonexist")
			break
		}
		err = nil
	}
	return err
}

func confFile() *conf {
	Warn.Println("making a default config")
	confDef := conf{
		OscIp:    "127.0.0.1",
		OscPort:  "8011",
		HttpPort: "80",
	}

	yamlData, _ := yaml.Marshal(confDef)
	err := os.WriteFile("config.yaml", yamlData, 0644)
	if err != nil {
		Error.Fatal("can't to write default conf into the file")
	}
	return &confDef
}

func ConfInit() *conf {
	cfg := &conf{}
	if err := getConf("config.yaml", cfg); err != nil {
		Error.Println(err)
		cfg = confFile()
	}

	if err := validateConf(cfg); err != nil {
		Error.Println(err)
		Error.Fatal("EXITING")
	}
	return cfg
}
