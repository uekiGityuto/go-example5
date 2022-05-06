package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/uekiGityuto/go-example5/handler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerJson(t *testing.T) {
	writer := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/fortune/json?name=Gopher", nil)
	handler.JSONHandler(writer, req)

	result := writer.Result()
	defer result.Body.Close()
	if result.StatusCode != http.StatusOK {
		t.Fatal("ステータスコードが200以外です")
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("レスポンスボディが取得出来ません")
	}
	var buffer bytes.Buffer
	buffer.Write(body)
	decoder := json.NewDecoder(&buffer)
	var fortune handler.Fortune
	if err := decoder.Decode(&fortune); err != nil {
		t.Fatal("レスポンスボディのデコードに失敗しました")
	}
	if fortune.Name != "Gopher" {
		t.Fatalf("レスポンスボディのnameが誤っています。expected: Gopher, actual: %s", fortune.Name)
	}
	switch fortune.Result {
	case "大吉":
	case "中吉":
	case "吉":
	case "小吉":
	case "末吉":
	case "凶":
	case "大凶":
	default:
		t.Fatalf("レスポンスボディのResultが誤っています。expected: [大吉, 中吉, 吉, 小吉, 末吉, 凶, 大凶], actual: %s", fortune.Result)
	}
}

func TestHandlerString(t *testing.T) {
	writer := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/fortune/string?name=Gopher", nil)
	handler.StringHandler(writer, req)

	result := writer.Result()
	defer result.Body.Close()
	if result.StatusCode != http.StatusOK {
		t.Fatal("ステータスコードが200以外です")
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("レスポンスボディが取得出来ません")
	}
	switch string(body) {
	case "Gopherさんの運勢は「大吉」です！\n":
	case "Gopherさんの運勢は「中吉」です！\n":
	case "Gopherさんの運勢は「吉」です！\n":
	case "Gopherさんの運勢は「小吉」です！\n":
	case "Gopherさんの運勢は「末吉」です！\n":
	case "Gopherさんの運勢は「凶」です！\n":
	case "Gopherさんの運勢は「大凶」です！\n":
	default:
		t.Fatalf("レスポンスボディが誤っています。expected: Gopherさんの運勢は「大吉|中吉|吉|小吉|末吉|凶|大凶」です！\n, actual: %s", string(body))
	}
}
