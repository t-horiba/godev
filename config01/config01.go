package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type config struct {
	Item1 string `json:"item1"`
	Item2 string `json:"item2"`
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
	fmt.Println(cfg.Item1)
	fmt.Println(cfg.Item2)

}
