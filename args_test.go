package translate

import (
	"testing"
)

func TestArgs(t *testing.T) {
	res := Args("name", "John")
	if len(res) != 1 {
		t.Error("Expected 1 result")
	} else if res["name"] != "John" {
		t.Error("Expected John result")
	}

	res = Args("name", "John", "age")
	if len(res) != 2 {
		t.Error("Expected 1 result")
	} else if res["name"] != "John" {
		t.Error("Expected John result")
	} else if res["age"] != nil {
		t.Error("Expected nil age")
	}
}
