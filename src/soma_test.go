package main

import "testing"

func TestSoma(t *testing.T) {
    total := soma(5, 5)
    if total != 10 {
       t.Errorf("Soma está incorreta, resultado: %d, esperado: %d.", total, 10)
    }
}
