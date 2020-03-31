package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	//dir = "./conf/"
	filename = "./conf/conf.json"
)

func init() {
	Conf = make(map[string]string)

	fmt.Println("filename:", filename)
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(filename + err.Error())
	}
	m := make(map[string]string)
	if err := json.Unmarshal(f, &m); err != nil {
		panic(filename + err.Error())
	}

	for key, value := range m {
		Conf.Set(key, value)
	}

	fmt.Printf("[ conf.json ] %#v\n", Conf)
}

type Config map[string]string

var Conf Config

func (c Config) Set(key, value string) {
	c[key] = value
}

func (c Config) Get(key string) (value string) {
	return c[key]
}
