package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Fortune struct {
	Name   string `json:"name"`
	Result string `json:"result"`
}

var results = []string{"大吉", "中吉", "吉", "小吉", "末吉", "凶", "大凶"}

func JSONHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var name string
	if req.FormValue("name") != "" {
		name = req.FormValue("name")
	} else {
		name = "名無し"
	}
	fortune := Fortune{Name: name, Result: results[rand.Intn(len(results))]}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	if err := encoder.Encode(fortune); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(writer, buffer.String())
}

func StringHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	var name string
	if req.FormValue("name") != "" {
		name = req.FormValue("name")
	} else {
		name = "名無し"
	}
	msg := name + "さんの運勢は「" + results[rand.Intn(len(results))] + "」です！"
	fmt.Fprintln(writer, msg)
}

func Listen() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/fortune/json", JSONHandler)
	http.HandleFunc("/fortune/string", StringHandler)

	http.ListenAndServe(":8080", nil)
}
