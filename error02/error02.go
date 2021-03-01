package main

// 参考：https://qiita.com/immrshc/items/13199f420ebaf0f0c37c
// go get github.com/pkg/errors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// ErrorType はエラーの種別を表す
type ErrorType uint

const (
	// Unknown は予期しないエラーを表す
	Unknown ErrorType = iota
	// InvalidArgument は引数エラーを表す
	InvalidArgument // = iota
	// Unauthorized は認証エラーを表す
	Unauthorized // = iota
	// ConnectionFailed は接続失敗を表す
	ConnectionFailed // = iota
)

type typeGetter interface {
	Type() ErrorType
}

type customError struct {
	errorType     ErrorType
	originalError error
}

// New 関数
func (et ErrorType) New(message string) error {
	return customError{errorType: et, originalError: errors.New(message)}
}

// Wrap 関数
func (et ErrorType) Wrap(err error, message string) error {
	return customError{errorType: et, originalError: errors.Wrap(err, message)}
}

// Error 関数
func (e customError) Error() string {
	return e.originalError.Error()
}

// Type 関数
func (e customError) Type() ErrorType {
	return e.errorType
}

// Wrap 関数
func Wrap(err error, message string) error {
	we := errors.Wrap(err, message)
	if ce, ok := err.(typeGetter); ok {
		return customError{errorType: ce.Type(), originalError: we}
	}
	return customError{errorType: Unknown, originalError: we}
}

// Cause 関数
func Cause(err error) error {
	return errors.Cause(err)
}

// GetType 関数
func GetType(err error) ErrorType {
	for {
		if e, ok := err.(typeGetter); ok {
			return e.Type()
		}
		break
	}
	return Unknown
}

// main 関数
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

// unmarshalToMap 関数
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
