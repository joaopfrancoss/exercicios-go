package main

import (
	"encoding/json"
	"exerciciogo/lorem"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const Largura = 100

type Contact struct {
	PrimeiroNome     string `json:"firstname"`
	Sobrenome        string `json:"lastname"`
	Numero           string `json:"phone"`
	Email            string `json:"email"`
	NumeroSeguidores int    `json:"followers"`
	Biografia        string
}

func leitor() {
	// pega a função open de os, para abrir o arquivo que esta no argumento, retornando o proprio arquivo e um erro.
	jsonFile, err := os.Open("contacts.json")
	// caso o erro aconteça, volta um panic
	if err != nil {
		panic(err)
	}
	// ultima coisa que a função vai fazer é rodar o defer, nesse caso, fechar o arquivo que abrimos.
	defer jsonFile.Close()

	// pegar a função readall de ioutil para ler os bites do arquivo json, retornando uma array de bites e erro.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	// variavel contacts vaia rmazenar uma array de contact (type contact struct...)
	var contacts []Contact
	// função unmarchal de json vai pegar os bites que ele ja leu e colocar no ponteiro de onde esta sendo armazenados os contatos da variavel contacts
	json.Unmarshal(byteValue, &contacts)

	for i := 0; i < len(contacts); i++ {
		contacts[i].Biografia = lorem.Paragraph(4, 6)
	}

	// para cada contact de contacts definido pelo i, ira pegar o primeiro nome e imprimir.
	indiceMinimoSeguidoresContato := 0
	indiceMaximoSeguidoresContato := 0
	for i := 0; i < len(contacts); i++ {
		if contacts[indiceMinimoSeguidoresContato].NumeroSeguidores > contacts[i].NumeroSeguidores {
			indiceMinimoSeguidoresContato = i
		}
		if contacts[indiceMaximoSeguidoresContato].NumeroSeguidores < contacts[i].NumeroSeguidores {
			indiceMaximoSeguidoresContato = i
		}

	}
	ficha(contacts[indiceMinimoSeguidoresContato])
	ficha(contacts[indiceMaximoSeguidoresContato])

	// somaSeguidores := 0
	// for i := 0; i < len(contacts); i++ {
	// 	somaSeguidores += contacts[i].NumeroSeguidores
	// }

	// media := somaSeguidores / len(contacts)
	// fmt.Println(media)

}

func emailCensurado(email string) string {
	var resultado string
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			resultado += email[i:]
			break
		} else {
			resultado += "x"
		}
	}
	return resultado
}

func ficha(contato Contact) {
	fmt.Println(strings.Repeat("-", Largura))
	nomeCompleto := contato.PrimeiroNome + " " + contato.Sobrenome
	re := regexp.MustCompile("[0-9]+")
	phoneRegexMatches := re.FindAllString(contato.Numero, -1)
	telefone := strings.Join(phoneRegexMatches, "")
	email := emailCensurado(contato.Email)
	biografia := contato.Biografia
	centralizado(nomeCompleto)
	centralizado(telefone)
	centralizado(email)
	quebraTexto(biografia)
	fmt.Println(strings.Repeat("-", Largura))
}

func centralizado(conteudo string) {
	var novoEspaco int
	k := (Largura - len(conteudo))
	numeroEspaco := k / 2
	novoEspaco = Largura - numeroEspaco - len(conteudo) - 1
	fmt.Println("|" + strings.Repeat(" ", numeroEspaco) + conteudo + strings.Repeat(" ", novoEspaco) + "|")
}

func quebraTexto(texto string) {
	var j int
	for i := 0; i < len(texto); i += 98 {
		j += 98
		if j > len(texto) {
			j = len(texto) - 1
		}
		for k := j; k > i; k-- {
			if texto[k] == ' ' {
				j = k
				break
			}
		}
		textoTamanho := "|" + texto[i:j]
		espacos := Largura - len(textoTamanho)
		fmt.Println(textoTamanho + strings.Repeat(" ", espacos) + "|")
	}

}
