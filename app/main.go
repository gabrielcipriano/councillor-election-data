package main

import "github.com/gabrielcipriano/sistema-eleitoral-vereadores/csvreader"

// func check(e error) {
// 	if e != nil {
// 		fmt.Println(e)
// 	}
// }

// Não é necessário produzir todos os 8 relatórios, apenas os relatórios
// “Número de vagas”, “Total de votos nominais” e mais 2 outros, à escolha;
// Os aparentemente mais faceis:
// • Candidatos eleitos (sempre indicado partido, número de votos e coligação, se houver)
// • Candidatos mais votados dentro do número de vagas;

func main() {
	csvreader.Read("../divulga20.csv")

}
