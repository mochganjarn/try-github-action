package main

import "testing"

func TestCetak(t *testing.T) {
	if cetak_hello() != "hello world" {
		t.Errorf("Salah Seharusnya hello world")
	}
}
