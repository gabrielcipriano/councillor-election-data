package main

import (
	"fmt"

	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/candidato"
	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/csvreader"
)

// Não é necessário produzir todos os 8 relatórios, apenas os relatórios
// “Número de vagas”, “Total de votos nominais” e mais 2 outros, à escolha;
// Os aparentemente mais faceis:
// • Candidatos eleitos (sempre indicado partido, número de votos e coligação, se houver)
// • Candidatos mais votados dentro do número de vagas;

func main() {
	var candidatos []candidato.Candidato
	candidatos = csvreader.Read("../divulga20.csv")
	for i, c := range candidatos {
		fmt.Printf("%d - %s\n", i+1, candidato.Stringify(&c))
	}

}
