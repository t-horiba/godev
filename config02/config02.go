package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type connect struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

type auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type config struct {
	Connect connect `json:"connect"`
	Auth    auth    `json:"auth"`
	Item1   string  `json:"item1"`
	Item2   string  `json:"item2"`
}

func main() {
	f, err := os.Open("./config.json")
	if err != nil {
		log.Fatal("loadConfig os.Open err:", err)
		return
	}
	defer f.Close()

	var cfg config

	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatal("NewDecode err:", err)
		return
	}
	fmt.Println(cfg.Connect.Hostname)
	fmt.Println(cfg.Connect.Port)
	fmt.Println(cfg.Auth.Username)
	fmt.Println(cfg.Auth.Password)
	fmt.Println(cfg.Item1)
	fmt.Println(cfg.Item2)
}
