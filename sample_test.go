package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name    string
	params  string
	returns string
}

func TestHello(t *testing.T) {
	result := Hello("ebi")
	if result != "Hello ebi" {
		t.Fatal("result must be 'Hello ebi'")
	}
}

func TestHelloAssert(t *testing.T) {
	result := Hello("ebi")
	assert.Equal(t, "Hello ebi", result, "result must be Hello ebi")
}

func TestHelloSubtest(t *testing.T) {
	t.Run("ebi params", func(t *testing.T) {
		result := Hello("ebi")
		assert.Equal(t, "Hello ebi", result, "result must be hello ebi")
	})

	t.Run("wahyu params", func(t *testing.T) {
		result := Hello("wahyu")
		assert.Equal(t, "Hello wahyu", result, "result must be hello ebi")
	})
}

func TestHelloTable(t *testing.T) {
	var test = []testTable{
		{
			name:    "test1",
			params:  "wahyu",
			returns: "Hello wahyu",
		},
		{
			name:    "test2",
			params:  "rizka",
			returns: "Hello rizka",
		},
	}

	for _, val := range test {
		t.Run(val.name, func(t *testing.T) {
			result := Hello(val.params)
			assert.Equal(t, val.returns, result)
		})
	}
}
