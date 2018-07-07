package main

import (
	"bufio"
	"github.com/standupdev/strset"
	"io"
	"strconv"
	"strings"
	"os"
	"log"
	"fmt"
)

type Resultado struct {
	codigo rune
	nome   string
}

func AnalisarLinha(texto string) Resultado {
	partes := strings.Split(texto, ";")
	codigo, _ := strconv.ParseInt(partes[0], 16, 32)
	runa := rune(codigo)
	return Resultado{runa, partes[1]}
}

func casarTermos(consulta string, nome string) bool {
	termosConsulta := strset.MakeFromText(
		strings.ToUpper(strings.Replace(consulta, "-", " ", -1)))
	termosNome := strset.MakeFromText(
		strings.ToUpper(strings.Replace(nome, "-", " ", -1)))
	return termosConsulta.SubsetOf(termosNome)
}

func BuscarRunas(consulta string, dados io.Reader) []Resultado {
	scanner := bufio.NewScanner(dados)
	achados := []Resultado{}
	for scanner.Scan() {
		res := AnalisarLinha(scanner.Text())
		if casarTermos(consulta, res.nome) {
			achados = append(achados, res)
		}
	}
	return achados
}

func main() {
	consulta := strings.Join(os.Args[1:], " ")
	ucd, err := os.Open("UnicodeData.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, res := range BuscarRunas(consulta, ucd) {
		fmt.Printf("%c\t%v\n", res.codigo, res.nome)
	}
}