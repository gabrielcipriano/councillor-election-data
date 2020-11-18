package main

import (
	// "encoding/csv"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/candidato"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// Não é necessário produzir todos os 8 relatórios, apenas os relatórios
// “Número de vagas”, “Total de votos nominais” e mais 2 outros, à escolha;
// Os aparentemente mais faceis:
// • Candidatos eleitos (sempre indicado partido, número de votos e coligação, se houver)
// • Candidatos mais votados dentro do número de vagas;

func main() {

	c := candidato.New(1, 1234, "Joao", "PT", "PT-PSOL", 12354, 3.333)
	fmt.Println(c)
	file, err := os.Open("divulga20.csv")
	check(err)

	reader := csv.NewReader(file)
	reader.Comma = ';'

	// line, err := reader.ReadAll()
	i := 0

	line, err := reader.Read() //descarta cabeçalho
	for {
		line, err = reader.Read()
		if err == io.EOF {
			break
		}
		check(err)
		fmt.Println(i, strings.HasPrefix(line[0], "*"))
		i++
	}

	defer file.Close()

}
