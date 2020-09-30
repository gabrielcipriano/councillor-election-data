package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type candidato struct {
	//situacao: 1 - candidato eleito; 0 - Normal; -1 - Inválido
	situacao          int
	numero            int
	nome              string
	partido           string
	coligacao         string
	votos             int
	percentualValidos float32
}

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
// • Votos totalizados por coligação ou partido (quando um partido não estiver em coligação), número de
// candidatos eleitos;
// • Votos totalizados por partido, número de candidatos eleitos;

func main() {
	// 	in := `first_name,last_name,username
	// "Rob","Pike",rob
	// Ken,Thompson,ken
	// "Robert","Griesemer","gri"
	// `
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
	// fmt.Println(line)

	defer file.Close()

	//r := csv.NewReader(strings.NewReader(in))
	// r := csv.NewReader(strings.NewReader(file))

	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(record)
	// }
}
