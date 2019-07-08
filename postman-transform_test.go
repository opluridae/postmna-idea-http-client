package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type PostamnData struct {
	Version      string         `json:"version"`
	Collections  []Collections  `json:"collections"`
	Environments []Environments `json:"environments"`
}

type Collections struct {
}
type Environments struct {
	Id     string      `json:"id"`
	Name   string      `json:"name"`
	Values []EnvValues `json:"values"`
}
type EnvValues struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Enable bool   `json:"enable"`
}

func Test_parsEnv(t *testing.T) {
	jsonFile, err := os.Open("C:/Users/tpavel/go/src/git.homebank.kz/mvisa_merchant/postman-data/Backup.postman_dump.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	data := PostamnData{}
	err = json.Unmarshal(byteValue, &data)
	for _, v := range data.Environments {
		fmt.Println("Env: " + v.Name)
		env := map[string]string{}

		for _, v := range v.Values {
			env[v.Key] = v.Value
		}

		bytes, err := json.MarshalIndent(env, "","  ")
		if err != nil {
			fmt.Println(err)
		}
		s :=  string(bytes)
		fmt.Println(s)
	}
}
