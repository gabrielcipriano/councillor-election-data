package candidato

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

//New candidato
func New(situacao int, numero int, nome string, partido string, coligacao string, votos int, percentualValidos float32) candidato {
	c := candidato{situacao, numero, nome, partido, coligacao, votos, percentualValidos}
	return c
}
