package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type MyXML struct{}

type Recurlyservers struct {
	XMLName     xml.Name    `xml:"servers"`
	Version     string      `xml:"version,attr"`
	Svs         []xmlserver `xml:"server"`
	Description string      `xml:",innerxml"`
}

type xmlserver struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func (my *MyXML) Read() {
	file, err := os.Open("servers.xml")
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
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	fmt.Println(v)
}

func (my *MyXML) Write() {
	v := &Recurlyservers{Version: "1"}
	v.Svs = append(v.Svs, xmlserver{ServerName: "Srv1", ServerIP: "localhost"})
	v.Svs = append(v.Svs, xmlserver{ServerName: "Srv2", ServerIP: "127.0.0.1"})
	output, err := xml.MarshalIndent(v, " ", "    ")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
