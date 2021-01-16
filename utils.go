package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

func p(a ...interface{}) {
	fmt.Println(a)
}

// init関数はパッケージの初期化に使われる
// mainパッケージにおくと、main関数より先に実行される
// https://qiita.com/tenntenn/items/7c70e3451ac783999b4f
func init() {
	loadConfig()

	// chitchat.logにログ吐き出す
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannnot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannnot get configuration from file", err)
	}
}

// エラーメッセージを表示するページに遷移させる
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func version() string {
	return "0.1"
}
