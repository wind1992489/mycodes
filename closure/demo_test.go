package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	counter := NewCounter()
	counter()
	counter()
	counter()
	counter()
}
