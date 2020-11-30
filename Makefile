build:
	$(info *** O comando make / make build cria o executável 'vereadores'. ***)
	$(info ***   execute da forma: ./vereadores caminho/para/arquivo.csv   ***)
	go build -o vereadores main.go

all:
	build


