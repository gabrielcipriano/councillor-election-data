package eleicao

import "github.com/gabrielcipriano/sistema-eleitoral-vereadores/candidato"

//Eleicao é a estrutura que guarda todas as informações pertinentes de uma eleicao
type Eleicao struct {
	NumEleitos            int //aka numero de vagas
	Candidatos            []candidato.Candidato
	TotalVotos            int
	candidatosMaisVotados []candidato.Candidato
}

// GetEleitos retorna a lista de candidatos eleitos
func (e *Eleicao) GetEleitos() []candidato.Candidato {
	return e.Candidatos[:e.NumEleitos]
}

func (e *Eleicao) computeMaisVotados() {
	e.candidatosMaisVotados = make([]candidato.Candidato, len(e.Candidatos))
	copy(e.candidatosMaisVotados, e.Candidatos)
	candidato.SortByVotes(e.candidatosMaisVotados)
}

// GetMaisVotados retorna a lista dos candidatos mais votados
func (e *Eleicao) GetMaisVotados() []candidato.Candidato {
	if len(e.candidatosMaisVotados) == 0 {
		e.computeMaisVotados()
	}
	tmp := make([]candidato.Candidato, len(e.candidatosMaisVotados))
	copy(tmp, e.candidatosMaisVotados)
	return tmp
}

// New retorna uma eleicão vazia
func New() Eleicao {
	var e Eleicao = Eleicao{
		NumEleitos:            0,
		TotalVotos:            0,
		Candidatos:            []candidato.Candidato{},
		candidatosMaisVotados: []candidato.Candidato{},
	}
	return e
}

// AddCandidato Adiciona um Candidato à eleição
func (e *Eleicao) AddCandidato(c candidato.Candidato) {
	e.Candidatos = append(e.Candidatos, c)
	e.TotalVotos += c.Votos
	if c.Situacao == candidato.Eleito {
		e.NumEleitos++
	}
}
