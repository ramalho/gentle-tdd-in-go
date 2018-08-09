package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {
	main()
	// Output:
	// Informe uma ou mais palavras para buscar.
}

func Example_1Arg() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "CRUZEIRO"}
	main()
	//Output:
	//U+20A2	₢	CRUZEIRO SIGN
}

func TestFormatar(t *testing.T) {
	testCases := []struct {
		nomeTeste string
		codigo    rune
		nome      string
		esperado  string
	}{
		{"teste A", 0x41, "LATIN CAPITAL LETTER A",
			"U+0041\tA\tLATIN CAPITAL LETTER A"},
		{"teste CRUZEIRO", 0x20a2, "CRUZEIRO SIGN",
			"U+20A2\t₢\tCRUZEIRO SIGN"},
	}
	for _, tc := range testCases {
		t.Run(tc.nomeTeste, func(t *testing.T) {
			resultado := formatar(tc.codigo, tc.nome)
			if resultado != tc.esperado {
				t.Errorf("Resultado: %s \t esperado: %s",
					resultado, tc.esperado)
			}
		})
	}
}

const dadosStr = `003C;LESS-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
003D;EQUALS SIGN;Sm;0;ON;;;;;N;;;;;
003E;GREATER-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
003F;QUESTION MARK;Po;0;ON;;;;;N;;;;;
0040;COMMERCIAL AT;Po;0;ON;;;;;N;;;;;
0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
0042;LATIN CAPITAL LETTER B;Lu;0;L;;;;;N;;;;0062;
`

func TestFiltrar(t *testing.T) {
	testCases := []struct {
		nomeTeste string
		palavras  []string
		esperado  []string
	}{
		{"?", []string{"QUESTION"},
			[]string{"003F;QUESTION MARK;Po;0;ON;;;;;N;;;;;"}},
		{"? minuscula", []string{"question"},
			[]string{"003F;QUESTION MARK;Po;0;ON;;;;;N;;;;;"}},
		{"inexistente", []string{"NAO EXISTE"},
			[]string{}},
		{"inexistente 2", []string{"MARK", "BIRTH"},
			[]string{}},
	}
	for _, tc := range testCases {
		t.Run(tc.nomeTeste, func(t *testing.T) {
			dados := strings.NewReader(dadosStr)
			resultado := filtrar(dados, tc.palavras)
			assert.Equal(t, tc.esperado, resultado)
		})
	}
}
