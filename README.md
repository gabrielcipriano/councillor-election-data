# sistema-eleitoral-vereadores
Ferramenta escrita em GO para gerar relatórios de eleições de vereadores.

## Instalação e execução:

#### Opção 1 - Com Makefile:
Baixe o arquivo .zip do repositório e extraia-o
##### Compilando:
```make```
ou
```make build```
##### Executando:
```./vereadores caminho/para/arquivo.csv```

#### Opção 2 - Clonando o repositório git diretamente para seu GOPATH:
##### Baixando:
```go get 'github.com/gabrielcipriano/sistema-eleitoral-vereadores'```
##### Instalando:
```go install github.com/gabrielcipriano/sistema-eleitoral-vereadores```
##### Executando:
O arquivo executável deve estar na pasta bin do seu GOPATH, isso é, `$USER/go/bin/`
Para executar:
```./sistema-eleitoral-vereadores caminho/para/arquivo.csv ```

##### Ao fim da execução, o arquivo `saida.txt` será criado no diretorio onde o programa foi executado.

#### Itens do relatório de saída:

- Total de vagas / total de candidatos eleitos;
- Total de votos nominais;

- Lista dos vereadores eleitos;

- Lista dos candidatos mais votados; (respeitando número de vagas)

- Lista dos candidatos que teriam sido eleitos se a votação fosse majoritária, e não foram eleitos;

- Lista dos candidatos eleitos que se beneficiaram do sistema proporcional;s