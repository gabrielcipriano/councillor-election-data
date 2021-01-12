# councillor-election-data
Program wrote in Golang to generate a report about local Brazilian Councillor election.

## Installation:

#### Option 1 - with Makefile:
Download the .zip from repository and extract it.
##### Compiling:
```make```
or
```make build```
##### Running:
```./vereadores path/to/file.csv```

#### Option 2 - Cloning the git repository directly to your GOPATH:
##### Download it:
```go get 'github.com/gabrielcipriano/councillor-election-data'```
##### Install it:
```go install github.com/gabrielcipriano/councillor-election-data```
##### Execute it:
The executation file should be in the `bin` folder of your GOPATH, that is, `$USER/go/bin/`
To run it:
```./councillor-election-data path/to/the/file.csv ```

##### At the end of the execution the file `saida.txt` (`exit.txt`) will be generated on the program folder.

#### Information in the report:

- Total vacancies / Total of ellected councillors;

- Total nominal votes;

- List of elected Councillors;

- List of most voted candidates; (respecting the number of vacancies)

- List of candidates who would be ellected if the election was majoritary, but wasn't.

- List of ellected candidates who took advantage on the proportional system.

### Entry files:
The **.csv** entry file are public data obteined by the tool 'DIVULGA' provided by the Superior Electoral Court.

On the folder `sample-csv-files` there is 2016 election data  from the cities of Rio de Janeiro, Vitoria and Vila Velha, as well as `expectedOutput.txt` files with the expected output.
