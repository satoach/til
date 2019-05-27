package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type MyJSON struct{}

type JSONServer struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Servers []JSONServer `json:"servers"`
}

func readJSON(f interface{}) {
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of type I don't how to handle")
		}
	}
}

func (my *MyJSON) Read() {
	file, err := os.Open("servers.json")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	var f interface{}
	err = json.Unmarshal(data, &f)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	readJSON(f)
}

func (my *MyJSON) Write() {
	var s Serverslice
	s.Servers = append(s.Servers, JSONServer{ServerName: "JSON1", ServerIP: "192.168.21.1"})
	s.Servers = append(s.Servers, JSONServer{ServerName: "JSON2", ServerIP: "127.0.0.1"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(string(b))
}
