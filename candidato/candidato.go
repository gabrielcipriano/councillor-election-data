package candidato

//Situacao enumera a situacao de um candidato: Eleito, Normal ou Invalido
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
	situacao          Situacao
	numero            int
	nome              string
	partido           string
	coligacao         string
	votos             int
	percentualValidos float32
}

//New Candidato from standard
func New(situacao Situacao, num int, nome string, partido string, colig string, votos int, percentualValidos float32) Candidato {
	c := Candidato{
		situacao:          situacao,
		numero:            num,
		nome:              nome,
		partido:           partido,
		coligacao:         colig,
		votos:             votos,
		percentualValidos: percentualValidos,
	}
	return c
}
