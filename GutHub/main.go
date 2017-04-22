package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"GutHub/srv/controller"
	fastpr "github.com/buaazp/fasthttprouter"
	fasthp "github.com/valyala/fasthttp"
)

const (
	initConfig = "./conf/init.json"
)

var config struct {
	Address string `json:"address"`
}

func init() {
	// read config file
	file, err := os.Open(initConfig)
	if err != nil {
		log.Fatalln("find init config file error:", err)
		os.Exit(-1)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("load init config file error:", err)
		os.Exit(-1)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln("unmarshal init config error:", err)
		os.Exit(-1)
	}

	if "" == config.Address {
		log.Fatalln("GutHub server not set.")
		os.Exit(-1)
	}
	log.Println("GutHub will start at:", config.Address)
}

func main() {
	log.Println("Welcome to GutHub.")

	router := fastpr.New()
	router.GET("/guts", controller.GetGutList)
	router.POST("/guts/add", controller.AddGut)
	router.PUT("/guts/modify/:id", controller.ModifyGut)
	router.DELETE("/guts/delete/:id", controller.DeleteGut)

	if err := fasthp.ListenAndServe(config.Address, router.Handler); err != nil {
		log.Println("start GutHub server failed, error:", err)
	}
}
