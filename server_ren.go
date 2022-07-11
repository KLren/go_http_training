package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type myStruct struct {
	Name string
}

type outPutStruct struct {
	Introduction string
}

type myStruct2 struct {
	Name string
}

func (m myStruct2) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	msg := "Hello go"

	switch req.Method {
	case "GET":
		fmt.Fprintln(w, msg)
	case "POST":
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Convert string to json format without struct
		var obj map[string]interface{}
		json.Unmarshal([]byte(b), &obj)
		msg = obj["name"].(string)
		fmt.Printf("handle2 - My name is %v\n", msg)
		fmt.Fprintf(w, "handle2 - My name is %v\n", msg)
	}
}

func myHandlerFunc1(w http.ResponseWriter, req *http.Request) {
	msg := "Hello go"

	switch req.Method {
	case "GET":
		fmt.Fprintln(w, msg)
	case "POST":
		var tm myStruct
		err := json.NewDecoder(req.Body).Decode(&tm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("de:%v , type: %T\n", tm, tm)
		sresp := outPutStruct{Introduction: "handle1 - My name is " + tm.Name}

		//ResponseWriter.Header – For writing response header
		//ResponseWriter.Write([]byte) – For writing response body
		//ResponseWriter.WriteHeader(statusCode int) – For writing the http status code
		jresp, err := json.Marshal(sresp)
		fmt.Printf("de:%s , type: %T\n", jresp, jresp)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jresp)
	}
}

func main() {
	var myhandler myStruct2

	http.HandleFunc("/myHandle1", myHandlerFunc1)
	http.Handle("/myHandle2", myhandler)
	http.ListenAndServe(":8080", nil)
}
