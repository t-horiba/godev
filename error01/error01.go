package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if _, err := unmarshalToMap("src.json"); err != nil {
		switch err2 := err.(type) {
		case *os.PathError:
			fmt.Printf("%+v\n", err2)
		case *json.SyntaxError:
			//fmt.Println("type: %s\n", err2.Type)
			fmt.Printf("offset: %d\n", err2.Offset)
			fmt.Printf("%+v\n", err2)
		default:
			fmt.Println(err2)
		}
	}
	if _, err := unmarshalToMap("dummy.json"); err != nil {
		switch err2 := err.(type) {
		case *os.PathError:
			fmt.Printf("%+v\n", err2)
		case *json.SyntaxError:
			//fmt.Println("type: %s\n", err2.Type)
			fmt.Printf("offset: %d\n", err2.Offset)
			fmt.Printf("%+v\n", err2)
		default:
			fmt.Println(err2)
		}
	}
}

func unmarshalToMap(src string) (map[string]interface{}, error) {
	jsonMap := map[string]interface{}{}
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return jsonMap, err
	}

	if err := json.Unmarshal(data, &jsonMap); err != nil {
		return nil, err
	}
	return jsonMap, nil
}
