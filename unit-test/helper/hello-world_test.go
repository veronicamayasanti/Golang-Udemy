package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldVeronica(t *testing.T) {
	result := HelloWorld("maya")

	if result != "Hello maya" {
		//panic("result is not 'Hello maya'")
		//t.Fail()
		t.Error("result must be 'Hello maya'")
	}
	fmt.Println("TestHelloWorldVeronica done", result)
}

func TestHelloWorldMaya(t *testing.T) {
	result := HelloWorld("maya")
	if result != "Hello maya" {
		//error
		//panic("result is not 'Hello maya'")
		//t.FailNow()
		t.Fatal("result must be 'Hello maya'")
	}
	fmt.Println("TestHelloWorldMaya done")
}
