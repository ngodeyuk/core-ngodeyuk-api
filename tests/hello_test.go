package tests

import (
	"ngodeyuk-core/pkg/utils"
	"testing"
)

func TestGetHelloMessage(t *testing.T) {
	expected := "Hello, World!"
	if msg := utils.GetHelloMessage(); msg != expected {
		t.Errorf("Expected %s but got %s", expected, msg)
	}
}
