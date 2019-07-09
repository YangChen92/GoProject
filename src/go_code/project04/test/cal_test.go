package main

import(
	// "fmt"
	"testing"
)

func TestCal(t *testing.T){
	res := addUpper(12)
	if res != 55 {
		t.Fatalf("addUpper error 期望值=%v 实际值=%v",55,res)
	}

	t.Logf("AddUpper 执行正确...")
}