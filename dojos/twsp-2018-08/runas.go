package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/standupdev/strset"
)

func busca(palavras []string) []string {
	dados, err := os.Open("UnicodeData.txt")
	if err != nil {
		panic(err)
	}
	defer func() { dados.Close() }()

	return filtrar(dados, palavras)
}

func formatar(codigo rune, nome string) string {
	return fmt.Sprintf("U+%04X\t%c\t%s", codigo, codigo, nome)
}

func bateu(linha string, palavras []string) bool {
	nome := strings.Split(linha, ";")[1]
	nomeSet := strset.MakeFromText(nome)
	palavrasSet := strset.MakeFromText(strings.ToUpper(strings.Join(palavras, " ")))
	return nomeSet.SupersetOf(palavrasSet)
}

func filtrar(dados io.Reader, palavras []string) []string {
	resultados := []string{}
	scanner := bufio.NewScanner(dados)
	for scanner.Scan() {
		linha := scanner.Text()
		if bateu(linha, palavras) {
			resultados = append(resultados, linha)
		}
	}
	return resultados
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Informe uma ou mais palavras para buscar.")

	} else {
		for _, linha := range busca(os.Args[1:]) {
			campos := strings.Split(linha, ";")
			campoCod := campos[0]
			codigo, _ := strconv.ParseInt(campoCod, 16, 32)
			nome := campos[1]
			fmt.Println(formatar(rune(codigo), nome))
		}
	}

}
