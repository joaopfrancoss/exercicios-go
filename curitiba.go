package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Solicitacao struct {
	NumeroSoliciacao int
	Tipo             string
	Orgao            string
	Data             string
	Horario          string
	Assunto          string
	Subdivisao       string
	Descricao        string
	Logradouro       string
	Bairro           string
}

func lerCsv() {
	csvfile, err := os.Open("baseDados.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	r.FieldsPerRecord = -1
	r.Comma = ';'

	var solicitacoes []Solicitacao
	var j int

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		j++
		if j == 1 || j == 2 {
			continue
		}
		if j > 80 {
			break
		}
		num, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}

		s := Solicitacao{
			NumeroSoliciacao: num,
			Tipo:             record[1],
			Orgao:            record[2],
			Data:             record[3],
			Horario:          record[4],
			Assunto:          record[5],
			Subdivisao:       record[6],
			Descricao:        record[7],
			Logradouro:       record[8],
			Bairro:           record[9],
		}

		solicitacoes = append(solicitacoes, s)
	}

	quantidade := 0
	for i := 0; i < len(solicitacoes); i++ {
		if solicitacoes[i].Bairro == "JARDIM BOTANICO" {
			quantidade++
		}
	}
	fmt.Println(quantidade)

	var ocorrenciaPorBairro = make(map[string]int)

	for i := 0; i < len(solicitacoes); i++ {
		_, existe := ocorrenciaPorBairro[solicitacoes[i].Bairro]
		if existe {
			ocorrenciaPorBairro[solicitacoes[i].Bairro]++
		} else {
			ocorrenciaPorBairro[solicitacoes[i].Bairro] = 1
		}
	}

	result, err := json.Marshal(ocorrenciaPorBairro)
	if err != nil {
		panic(err)
	}

	resultStr := string(result)

	jj, err := os.Create("index.json")
	if err != nil {
		panic(err)
	}
	defer jj.Close()

	_, err = jj.WriteString(resultStr)
	if err != nil {
		panic(err)
	}
}
