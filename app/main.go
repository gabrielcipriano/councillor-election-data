package main

import (
	"fmt"

	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/candidato"
	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/csvreader"
	elec "github.com/gabrielcipriano/sistema-eleitoral-vereadores/eleicao"
)

// Não é necessário produzir todos os 8 relatórios, apenas os relatórios
// “Número de vagas”, “Total de votos nominais” e mais 2 outros, à escolha;
// Os aparentemente mais faceis:
// • Candidatos eleitos (sempre indicado partido, número de votos e coligação, se houver)
// • Candidatos mais votados dentro do número de vagas;

func imprimeNumDeVagas(eleicao *elec.Eleicao) {
	fmt.Printf("Número de vagas: %d\n\n", eleicao.NumEleitos)
}
func imprimeTotalVotos(eleicao *elec.Eleicao) {
	fmt.Printf("Total de votos nominais: %d\n\n", eleicao.TotalVotos)
}
func imprimeCandidatos(candidatos []candidato.Candidato, len int) {
	for i := 0; i < len; i++ {
		fmt.Printf("%d - %s\n", i+1, candidatos[i].ToString())
	}
}

//Imprime somente candidatos que obedançam a uma determinada condição
func imprimeCandidatosPrejudicados(eleicao *elec.Eleicao) {
	maisVotados := eleicao.GetMaisVotados()
	eleitoMenosVotado := eleicao.Candidatos[eleicao.NumEleitos-1]
	for i := 0; i < eleicao.NumEleitos; i++ {
		eleito := maisVotados[i].Situacao == candidato.Eleito
		votos := maisVotados[i].Votos
		if !eleito && votos > eleitoMenosVotado.Votos {
			fmt.Printf("%d - %s\n", i+1, maisVotados[i].ToString())
		}
	}
}

func imprimeCandidatosBeneficiados(eleicao *elec.Eleicao) {
	maisVotados := eleicao.GetMaisVotados()
	numEleitos := eleicao.NumEleitos
	naoSeriamEleitos := maisVotados[numEleitos:]
	for i := 0; i < len(naoSeriamEleitos); i++ {
		if naoSeriamEleitos[i].Situacao == candidato.Eleito {
			fmt.Printf("%d - %s\n", i+numEleitos+1, naoSeriamEleitos[i].ToString())
		}
	}
}

func main() {
	//TODO: Receber nome do arquivo como paramentro do programa

	eleicao := elec.New()
	csvreader.Read("../divulga.csv", &eleicao)
	// for i, cand := range eleicao.Candidatos {
	// 	fmt.Printf("%d - %s\n", i+1, cand.ToString())
	// }

	imprimeNumDeVagas(&eleicao)

	imprimeTotalVotos(&eleicao)

	fmt.Println("Vereadores eleitos:")
	imprimeCandidatos(eleicao.GetEleitos(), eleicao.NumEleitos)

	//Relatório "Candidatos mais votados dentro do numero de vagas"
	fmt.Println("\nCandidatos mais votados (em ordem decrescente de votação e respeitando número de vagas):")
	imprimeCandidatos(eleicao.GetMaisVotados(), eleicao.NumEleitos)

	//BONUS: Candidatos não eleitos e que seriam eleitos se a votação fosse majoritária;
	fmt.Println("\nTeriam sido eleitos se a votação fosse majoritária, e não foram eleitos:")
	fmt.Println("(com sua posição no ranking de mais votados)")
	imprimeCandidatosPrejudicados(&eleicao)

	//BONUS: Eleitos, que se beneficiaram do sistema proporcional:
	fmt.Println("\nEleitos, que se beneficiaram do sistema proporcional:")
	fmt.Println("(com sua posição no ranking de mais votados)")
	imprimeCandidatosBeneficiados(&eleicao)

	//TODO: Gravar relatórios no arquivo "saida.txt"

}
