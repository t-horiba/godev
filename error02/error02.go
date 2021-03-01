package main

// 参考：https://qiita.com/immrshc/items/13199f420ebaf0f0c37c
// go get github.com/pkg/errors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// ErrorType はエラーの種別を表し、New メソッド、Wrap メソッドを含む
type ErrorType uint

// New メソッド
func (et ErrorType) New(message string) error {
	return customError{errorType: et, originalError: errors.New(message)}
}

// Wrap メソッド
func (et ErrorType) Wrap(err error, message string) error {
	return customError{errorType: et, originalError: errors.Wrap(err, message)}
}

const (
	// Unknown は予期しないエラーを表す
	Unknown ErrorType = iota
	// InvalidArgument は引数エラーを表す
	InvalidArgument // ErrorType = iota
	// Unauthorized は認証エラーを表す
	Unauthorized // ErrorType = iota
	// ConnectionFailed は接続失敗を表す
	ConnectionFailed // ErrorType = iota
)

// typeGetter インターフェースの宣言
type typeGetter interface {
	// メソッドリスト
	Type() ErrorType
}

// customError 構造体の宣言
type customError struct {
	// フィールドリスト
	errorType ErrorType
	// Error() メソッドを含む構造体を代入
	originalError error
	// Error() メソッド
	// Type() メソッド
}

// Error メソッド
func (e customError) Error() string {
	return e.originalError.Error()
}

// Type メソッド
func (e customError) Type() ErrorType {
	return e.errorType
}

// Wrap 関数
//func Wrap(err error, message string) error {
//	we := errors.Wrap(err, message)
//	if ce, ok := err.(typeGetter); ok {
//		return customError{errorType: ce.Type(), originalError: we}
//	}
//	return customError{errorType: Unknown, originalError: we}
//}

// Cause 関数
//func Cause(err error) error {
//	return errors.Cause(err)
//}

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
	err := Unauthorized.New("認証エラー")
	fmt.Println(statusCode(err))

	err = panicAndRecover()
	fmt.Println(err)

	if _, err1 := unmarshalToMap("src.json"); err1 != nil {
		err1 = InvalidArgument.Wrap(err1, "引数エラー")
		fmt.Println(statusCode(err1))
		fmt.Println(err1.Error())
		fmt.Printf("%+v\n", err1)
	}
}

func statusCode(err error) int {
	switch GetType(err) {
	case ConnectionFailed:
		return http.StatusInternalServerError
	case Unauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusBadRequest
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

func main2() {
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
