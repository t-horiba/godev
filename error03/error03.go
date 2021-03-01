package main

// 参考：https://qiita.com/immrshc/items/13199f420ebaf0f0c37c
// go get github.com/pkg/errors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

// customError 構造体の宣言
type customError struct {
	// フィールドリスト
	field1        string
	field2        string
	field3        string
	originalError error
	// Error() メソッド
	// New() メソッド
	// Wrap() メソッド
}

// Error メソッド
func (e customError) Error() string {
	return e.originalError.Error()
}

// New メソッド
func (e customError) New(value1 string, value2 string, value3 string) error {
	return customError{field1: value1, field2: value2, field3: value3, originalError: nil}
}

// Wrap メソッド
func (e customError) Wrap(err error, value1 string, value2 string, value3 string) error {
	return customError{field1: value1, field2: value2, field3: value3, originalError: err}
}

// main 関数
func main() {
	if _, err1 := unmarshalToMap("src.json"); err1 != nil {
		fmt.Println(err1.field1)
		fmt.Println(err1.field2)
		fmt.Println(err1.field3)
		fmt.Println(err1.Error())
		fmt.Printf("%+v\n", err1.Error())
	}
}

func panicAndRecover() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recovered: %v\n", r))
		}
	}()
	panic("panic at panicAndRecover")
	// return
}

// unmarshalToMap 関数
func unmarshalToMap(src string) (map[string]interface{}, customError) {
	jsonMap := map[string]interface{}{}
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return jsonMap, customError.Wrap(err, "unmarchalToMap", "ioutil.ReadFile", src)
		//return jsonMap, err
	}

	if err := json.Unmarshal(data, &jsonMap); err != nil {
		return nil, customError.Wrap(err, "unmarchalToMap", "json.Unmarshal", "")
		//return nil, err
	}
	return jsonMap, nil
}
