package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := soma(40, 2)
	if resultado != 42 {
		t.Errorf("Soma incorreta. Esperado %v, encontrado %v", 42, resultado)
	}
}