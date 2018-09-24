package main

import (
	"fmt"
	"time"
	"k8s.io/client-go/util/jsonpath"
	"bytes"
	"encoding/json"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {
	//const TimeFormat = "2006-01-02T15:04:05"
	//t,_:=time.Parse(TimeFormat,"2013-08-11T11:18:46")
	//fmt.Print(t)
	var input = []byte(`{
	  "kind": "List",
	  "items":[
		{
		  "kind":"None",
		  "metadata":{
		    "name":"127.0.0.1",
			"labels":{
			  "kubernetes.io/hostname":"127.0.0.1"
			}
		  },
		  "status":{
			"capacity":{"cpu":"4"},
			"ready": true,
			"addresses":[{"type": "LegacyHostIP", "address":"127.0.0.1"}]
		  }
		},
		{
		  "kind":"None",
		  "metadata":{
			"name":"127.0.0.2",
			"labels":{
			  "kubernetes.io/hostname":"127.0.0.2"
			}
		  },
		  "status":{
			"capacity":{"cpu":"8"},
			"ready": false,
			"addresses":[
			  {"type": "LegacyHostIP", "address":"127.0.0.2"},
			  {"type": "another", "address":"127.0.0.3"}
			]
		  }
		}
	  ],
	  "users":[
	    {
	      "name": "myself",
	      "user": {}
	    },
	    {
	      "name": "e2e",
	      "user": {"username": "admin", "password": "secret"}
	  	}
	  ]
	}`)
	var data interface{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		fmt.Println(err.Error())
	}
	jp := jsonpath.New("test")
	jp.AllowMissingKeys(false)
	err1 := jp.Parse("aaa")
	if err1 != nil {
		fmt.Println(err1)
	}
	buf := new(bytes.Buffer)
	err2 := jp.Execute(buf, data)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	out := buf.String()
	fmt.Println(out)
}

