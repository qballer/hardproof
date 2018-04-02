package main 

import (
	"testing"
)

func TestAdd(t *testing.T){
	result := Add(4,  5);
	if result != 9 {
		t.Errorf("Add doesnt work, result is: " + string(result))
	}
}