package main

import (
	"testing"
	"strings"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestAnalisarLinha(t *testing.T) {
	linha := "0028;LEFT PARENTHESIS;Ps;0;ON;;;;;Y;OPENING PARENTHESIS;;;;"
	resultado := AnalisarLinha(linha)
	esperado := Resultado{'(', "LEFT PARENTHESIS"}

	if resultado.codigo != esperado.codigo {
		t.Errorf("Runa incorreta: %q, esperado: %q", resultado.codigo, esperado.codigo)
	}

	if resultado.nome != esperado.nome {
		t.Errorf("Nome incorreto: %#v, esperado: %#v", resultado.nome, esperado.nome)
	}
}

const dados = `0024;DOLLAR SIGN;Sc;0;ET;;;;;N;;;;;
0025;PERCENT SIGN;Po;0;ET;;;;;N;;;;;
0026;AMPERSAND;Po;0;ON;;;;;N;;;;;
0027;APOSTROPHE;Po;0;ON;;;;;N;APOSTROPHE-QUOTE;;;;
`

func TestBuscarRunas(t *testing.T) {

	resultados := BuscarRunas("sign", strings.NewReader(dados))
	esperado := []Resultado{{'$', "DOLLAR SIGN"}, {'%', "PERCENT SIGN"}}

	assert.Equal(t, esperado, resultados, "diferentes")
}

func TestBuscarRunas_inexistente(t *testing.T) {
	resultados := BuscarRunas("nãoexisteisso", strings.NewReader(dados))
	esperado := []Resultado{}
	assert.Equal(t, esperado, resultados, "diferentes")
}

func TestCasarTermos(t *testing.T) {
	testCases := []struct {
		consulta  	string
		nome  		string
		esperado	bool
	}{
		{"A", "A", true},
		{"A", "AB", false},
		{"a", "A B", true},
		{"B", "A-B", true},
		{"B-A", "A-B C", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s em %s", tc.consulta, tc.nome), func(t *testing.T) {
			res := casarTermos(tc.consulta, tc.nome)
			if res != tc.esperado {
				t.Errorf("veio %v; esperado %v", res, tc.esperado)
			}
		})
	}
}

func Example() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "registered"}
	main()
	// Output:
	// ®	REGISTERED SIGN
}
