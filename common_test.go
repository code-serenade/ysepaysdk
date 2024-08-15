package ysepaysdk

import (
	"fmt"
	"testing"
)

var Shared Config

func init() {
	Shared, _ = NewConfigFromFile("temp.json")
	Verbose = true
}

func TestXxx(t *testing.T) {
	req := NewAPIRequest("", "", "", "")
	fmt.Printf("%#v", req)
}
