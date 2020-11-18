package csvreader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gabrielcipriano/sistema-eleitoral-vereadores/candidato"
)

const (
	// prefixRegex string = `([*#])`
	// numeroRegex   string = `(\d+)`  acho que nao precisa

	// separador dos csv's entregues pelo TSE: ";"
	separator = ';'
)

// regex que identifica o prefixo da linha, caso exista ( * ou # )
// var /* const */ prefixPattern = regexp.MustCompile(prefixRegex)

// var /* const */ numeroPattern = regexp.MustCompile(numeroRegex) //acho que nao precisa

const (
	situacaoField = iota
	numField
	candidatoField
	partidoEColigField
	numVotosField
)

//CsvReader estrutura da leitura de arquivos CSV
type CsvReader struct {
	file       *os.File
	reader     csv.Reader
	candidatos []candidato.Candidato
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

//descarta o cabeçalho
func descartaCabecalho(reader *csv.Reader) {
	_, err := reader.Read()
	check(err)
}

//returns the Situacao from a prefix, that means, "*" = Eleito; "#" = Invalido, otherwise: Normal
func situacaoFromField(situacaoField string) candidato.Situacao {
	switch {
	case strings.HasPrefix(situacaoField, "*"):
		return candidato.Eleito
	case strings.HasPrefix(situacaoField, "#"):
		return candidato.Invalido
	default:
		return candidato.Normal
	}
}

//returns a Candidato from a splited csv line
func candidatoFromLine(line []string) candidato.Candidato {
	c := candidato.Candidato{}

	return c
}

//Read reads the csv file and returns a list of Candidatos in the file
func Read(filePath string) /*[]candidato.Candidato */ {
	file, err := os.Open(filePath)
	check(err)

	reader := csv.NewReader(file)
	reader.Comma = separator

	descartaCabecalho(reader)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		check(err)
		// prefixos possiveis: "*" ou "#" ou ""
		prefix := prefixPattern.FindString(line[situacaoField])

		fmt.Printf("%s %T\n", prefix, prefix)
	}

	defer file.Close()

}
