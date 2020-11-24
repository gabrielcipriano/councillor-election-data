package candidato

import (
	"sort"
	"strconv"
)

// Sorting Candidatos by number of votes

type byVotes []Candidato

func (candidatos byVotes) Len() int {
	return len(candidatos)
}
func (candidatos byVotes) Swap(i, j int) {
	candidatos[i], candidatos[j] = candidatos[j], candidatos[i]
}
func (candidatos byVotes) Less(i, j int) bool {
	return candidatos[i].Votos > candidatos[j].Votos
}

//SortByVotes Ordena uma lista de Candidatos por número de votos
func SortByVotes(candidatos []Candidato) []Candidato {
	sort.Sort(byVotes(candidatos))
	return candidatos
}

//Situacao enumera a situacao de um candidato: 0 = Eleito, 1 = Normal ou 2 = Invalido
type Situacao int

const (
	//Eleito - candidato eleito
	Eleito Situacao = iota
	//Normal - candidato regular que não foi eleito
	Normal
	//Invalido - candidato invalido por algum motivo
	Invalido
)

func (s Situacao) toString() string {
	situacoes := [...]string{
		"Candidato eleito",
		"Normal",
		"Inválido",
	}
	return situacoes[s]
}

//Candidato é a estrutura que guarda todas as informações pertinentes de um candidato
type Candidato struct {
	Situacao  Situacao
	numero    int
	Nome      string
	partido   string
	coligacao string
	Votos     int
}

// ToString retorna uma versão printável da struct Candidato
func (c *Candidato) ToString() string {
	s := c.Nome + " (" + c.partido + ", " + strconv.Itoa(c.Votos) + ")"
	if c.coligacao != "" {
		s = s + " - Coligação: " + c.coligacao
	}
	return s
}

//New Candidato
func New(situacao Situacao, num int, nome string, partido string, colig string, votos int) Candidato {
	c := Candidato{
		Situacao:  situacao,
		numero:    num,
		Nome:      nome,
		partido:   partido,
		coligacao: colig,
		Votos:     votos,
	}
	return c
}
