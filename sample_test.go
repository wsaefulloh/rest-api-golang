package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("ebi")
	if result != "Hello ebi" {
		t.Fatal("result must be 'Hello ebi'")
	}
}
