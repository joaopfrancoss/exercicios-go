package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	var template = `
	<div class= "ep">
		</br> 
		%s 
		<span class = "espaço"></span> 
		<button> 
			<a href= %s   target="_blank" >Link</a>  
		</button>
		 %s 
		 <img src= %s>
		 </br>
	 	</br>
	 </div>
	 </br>
	 </br>
	 `

	var html = `
	<html>
	<style>
	.ep{
		background-color: #20B2AA;
		text-align: center;
	}

	.espaço{
		width: 15px;
	}
	</style>
	<h1> oi </h1>
	`
	f, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	narcos := lerjson()

	var epsPorTemporada = make(map[int][]Episodio)

	for i := 0; i < len(narcos.Embedded.Episodios); i++ {
		_, existe := epsPorTemporada[narcos.Embedded.Episodios[i].Temporada]
		if existe {
			epsPorTemporada[narcos.Embedded.Episodios[i].Temporada] = append(epsPorTemporada[narcos.Embedded.Episodios[i].Temporada],
				narcos.Embedded.Episodios[i])
		} else {
			epsPorTemporada[narcos.Embedded.Episodios[i].Temporada] = []Episodio{}
		}
	}

	for temporada, episodios := range epsPorTemporada {
		temp := strconv.Itoa(temporada)
		html += "<h1>" + "TEMPORADA " + temp + "</h1>" + "</br>"
		for i := 0; i < len(episodios); i++ {
			html += fmt.Sprintf(template, episodios[i].Nome, episodios[i].Link, episodios[i].Comentario, episodios[i].Imagens.Medio)

		}
	}

	// for i := 0; i < len(narcos.Embedded.Episodios); i++ {
	// 	if narcos.Embedded.Episodios[i].Temporada == 1 {
	// 		html += narcos.Embedded.Episodios[i].Nome + "<img src=" + narcos.Embedded.Episodios[i].Imagens.Medio + ">" + "</br>"
	// 	}
	// }

	html += "</html>"
	_, err = f.WriteString(html)
	if err != nil {
		panic(err)
	}

}

type Narcos struct {
	Embedded Embedded `json:"_embedded"`
}

type Embedded struct {
	Episodios []Episodio `json:"episodes"`
}

type Episodio struct {
	Identificador int              `json:"id"`
	Link          string           `json:"url"`
	Nome          string           `json:"name"`
	Temporada     int              `json:"season"`
	Imagens       EpisodiosImagens `json:"image"`
	Comentario    string           `json:"summary"`
}

type EpisodiosImagens struct {
	Medio string `json:"medium"`
}

func lerjson() Narcos {
	jsonFile, err := os.Open("narcos.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var narcos Narcos
	json.Unmarshal(byteValue, &narcos)
	return narcos
}

func nomeEp(ep Episodio) {

}
