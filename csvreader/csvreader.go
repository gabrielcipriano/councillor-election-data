package csvreader

import (
	"encoding/csv"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/candidato"
	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/utils"
)

const (
	// separador dos csv's entregues pelo TSE: ";"
	separator = ';'
)

// No csv o partido é separado da coligação por um " - "
var /* const */ separadorPartidoColigPattern = regexp.MustCompile(` - `)

//

//constantes para nomear os indexes das colunas do arquivo, indo de 0 a 5
const (
	situacaoField = iota
	numField
	nomeField
	partidoEColigField
	numVotosField
)

//descarta o cabeçalho do arquivo CSV
func descartaCabecalho(reader *csv.Reader) {
	_, err := reader.Read()
	utils.CheckError(err)
}

// Retorna a situação de um candidato de acordo com o prefixo, isso é, "*": Eleito; "#": Invalido; Qualquer outro prefixo: Normal
func getSituacaoFromField(situacaoField string) candidato.Situacao {
	switch {
	case strings.HasPrefix(situacaoField, "*"):
		return candidato.Eleito
	case strings.HasPrefix(situacaoField, "#"):
		return candidato.Invalido
	default:
		return candidato.Normal
	}
}

func getNumVotosFromField(stringNumVotos string) int {
	var numVotos int = utils.IntFromHumanRedableNumber(stringNumVotos)
	return numVotos
}

func getPartidoEColicagaoFromField(partidoEColigField string) (partido string, coligacao string) {
	// Caso exista coligacao, parte a string do partido e coligacao em duas, retornando uma lista com as duas strings.
	var partidoEColigacao []string = separadorPartidoColigPattern.Split(partidoEColigField, 2)
	partido = partidoEColigacao[0]

	// adicionando coligacao caso exista
	if len(partidoEColigacao) > 1 {
		coligacao = partidoEColigacao[1]
	} else {
		coligacao = ""
	}
	return
}

//Retorna um Candidato dada uma linha do CSV
func candidatoFromLine(line []string) candidato.Candidato {
	var situacao candidato.Situacao
	var numero, numVotos int
	var nome, partido, coligacao string

	situacao = getSituacaoFromField(line[situacaoField])
	numero, _ = strconv.Atoi(line[numField])
	numVotos = getNumVotosFromField(line[numVotosField])
	nome = line[nomeField]
	partido, coligacao = getPartidoEColicagaoFromField(line[partidoEColigField])

	candidato := candidato.New(situacao, numero, nome, partido, coligacao, numVotos)

	return candidato
}

func candidatosFromReaderLines(reader *csv.Reader) []candidato.Candidato {
	var candidatos []candidato.Candidato
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		utils.CheckError(err)

		c := candidatoFromLine(line)

		candidatos = append(candidatos, c)
	}

	return candidatos
}

//Read lê o arquivo CSV do parâmetro e retorna uma lista dos objetos 'Candidato' nesse arquivo.
func Read(filePath string) []candidato.Candidato {
	file, err := os.Open(filePath)
	utils.CheckError(err)

	reader := csv.NewReader(file)
	reader.Comma = separator

	descartaCabecalho(reader)

	candidatos := candidatosFromReaderLines(reader)

	defer file.Close()

	return candidatos
}
