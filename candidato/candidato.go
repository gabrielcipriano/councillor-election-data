package candidato

import "strconv"

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
	situacao  Situacao
	numero    int
	nome      string
	partido   string
	coligacao string
	votos     int
}

// Stringify retorna uma versão printável da struct Candidato
func Stringify(c *Candidato) string {
	s := c.nome + " (" + c.partido + ", " + strconv.Itoa(c.votos) + ") - "
	s = s + "Coligação: " + c.coligacao + " - Situação: " + c.situacao.toString()
	return s
}

//New Candidato
func New(situacao Situacao, num int, nome string, partido string, colig string, votos int) Candidato {
	c := Candidato{
		situacao:  situacao,
		numero:    num,
		nome:      nome,
		partido:   partido,
		coligacao: colig,
		votos:     votos,
	}
	return c
}
