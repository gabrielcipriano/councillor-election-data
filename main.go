package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/candidato"
	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/csvreader"
	elec "github.com/gabrielcipriano/sistema-eleitoral-vereadores/eleicao"
	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/utils"
)

func escreveNumDeVagas(w *bufio.Writer, eleicao *elec.Eleicao) {
	fmt.Fprintf(w, "Número de vagas: %d\n\n", eleicao.NumEleitos)
}
func escreveTotalVotos(w *bufio.Writer, eleicao *elec.Eleicao) {
	fmt.Fprintf(w, "Total de votos nominais: %d\n\n", eleicao.TotalVotos)
}
func escreveCandidatos(w *bufio.Writer, candidatos []candidato.Candidato, len int) {
	for i := 0; i < len; i++ {
		fmt.Fprintf(w, "%d - %s\n", i+1, candidatos[i].ToString())
	}
}

//Imprime somente candidatos que foram prejudicados pelo sistema eleitoral
func escreveCandidatosPrejudicados(w *bufio.Writer, eleicao *elec.Eleicao) {
	maisVotados := eleicao.GetMaisVotados()
	eleitoMenosVotado := eleicao.Candidatos[eleicao.NumEleitos-1]
	for i := 0; i < eleicao.NumEleitos; i++ {
		eleito := maisVotados[i].Situacao == candidato.Eleito
		votos := maisVotados[i].Votos
		if !eleito && votos > eleitoMenosVotado.Votos {
			fmt.Fprintf(w, "%d - %s\n", i+1, maisVotados[i].ToString())
		}
	}
}

//Imprime somente candidatos que foram beneficiados pelo sistema eleitoral
func escreveCandidatosBeneficiados(w *bufio.Writer, eleicao *elec.Eleicao) {
	maisVotados := eleicao.GetMaisVotados()
	numVagas := eleicao.NumEleitos
	naoSeriamEleitos := maisVotados[numVagas:]
	for i := 0; i < len(naoSeriamEleitos); i++ {
		if naoSeriamEleitos[i].Situacao == candidato.Eleito {
			fmt.Fprintf(w, "%d - %s\n", i+numVagas+1, naoSeriamEleitos[i].ToString())
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("arquivo como argumento está faltando.")
		fmt.Printf("execute da forma: %s caminho/para/arquivo.csv", os.Args[0])
		os.Exit(1)
	}
	filepath := os.Args[1]

	eleicao := elec.New()
	csvreader.Read(filepath, &eleicao)

	//Imprimindo o relatorio no arquivo saida.txt
	file, err := os.Create("saida.txt")
	utils.CheckError(err)
	defer file.Close()

	w := bufio.NewWriter(file)

	escreveNumDeVagas(w, &eleicao)

	escreveTotalVotos(w, &eleicao)

	w.WriteString("Vereadores eleitos:\n")
	escreveCandidatos(w, eleicao.GetEleitos(), eleicao.NumEleitos)

	//Relatório "Candidatos mais votados dentro do numero de vagas"
	w.WriteString("\nCandidatos mais votados (em ordem decrescente de votação e respeitando número de vagas):\n")
	escreveCandidatos(w, eleicao.GetMaisVotados(), eleicao.NumEleitos)

	//BONUS: Candidatos não eleitos e que seriam eleitos se a votação fosse majoritária;
	w.WriteString("\nTeriam sido eleitos se a votação fosse majoritária, e não foram eleitos:\n")
	w.WriteString("(com sua posição no ranking de mais votados)\n")
	escreveCandidatosPrejudicados(w, &eleicao)

	//BONUS: Eleitos, que se beneficiaram do sistema proporcional:
	w.WriteString("\nEleitos, que se beneficiaram do sistema proporcional:\n")
	w.WriteString("(com sua posição no ranking de mais votados)\n")
	escreveCandidatosBeneficiados(w, &eleicao)

	w.Flush()

	fmt.Println("Arquivo \"saida.txt\" criado/sobrescrito")

}
